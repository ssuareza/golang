package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

var (
	bucket    = "file-random-322"
	object    = "file_500MB.bin"
	localFile = "./file_500MB.bin"
)

func main() {
	// Initialize GCS client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to create client: %w", err))
	}
	defer client.Close()

	// Get object attributes to determine size
	attrs, err := client.Bucket(bucket).Object(object).Attrs(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to get object attrs: %w", err))
	}
	size := attrs.Size

	// Create local file with correct size
	f, err := os.Create(localFile)
	if err != nil {
		panic(fmt.Errorf("failed to create local file: %w", err))
	}
	defer f.Close()
	if err := f.Truncate(size); err != nil {
		panic(fmt.Errorf("failed to set file size: %w", err))
	}

	// Parallel chunk download
	const chunkSize int64 = 12 * 1024 * 1024 // 8MB per chunk
	numChunks := int((size + chunkSize - 1) / chunkSize)
	maxParallel := 12
	type chunk struct {
		idx   int
		start int64
		end   int64
	}

	chunks := make([]chunk, numChunks)
	for i := 0; i < numChunks; i++ {
		start := int64(i) * chunkSize
		end := start + chunkSize
		if end > size {
			end = size
		}
		chunks[i] = chunk{idx: i, start: start, end: end}
	}

	startTime := time.Now()
	errCh := make(chan error, numChunks)
	doneCh := make(chan struct{}, numChunks)

	sem := make(chan struct{}, maxParallel)

	for _, c := range chunks {
		sem <- struct{}{} // acquire semaphore
		go func(c chunk) {
			defer func() { <-sem }() // release semaphore
			r, err := client.Bucket(bucket).Object(object).NewRangeReader(ctx, c.start, c.end-c.start)
			if err != nil {
				errCh <- fmt.Errorf("chunk %d: failed to create range reader: %w", c.idx, err)
				return
			}
			defer r.Close()

			buf := make([]byte, c.end-c.start)
			if _, err := io.ReadFull(r, buf); err != nil {
				errCh <- fmt.Errorf("chunk %d: failed to read: %w", c.idx, err)
				return
			}

			// Write to file at correct offset
			if _, err := f.WriteAt(buf, c.start); err != nil {
				errCh <- fmt.Errorf("chunk %d: failed to write: %w", c.idx, err)
				return
			}

			doneCh <- struct{}{}
		}(c)
	}

	// Wait for all chunks
	completed := 0
	var firstErr error
	for completed < numChunks {
		select {
		case <-doneCh:
			completed++
		case err := <-errCh:
			if firstErr == nil {
				firstErr = err
			}
			completed++
		}
	}

	duration := time.Since(startTime)

	if firstErr != nil {
		panic(firstErr)
	}

	fmt.Printf("Downloaded gs://%s/%s to %s in %s using %d parallel chunks\n", bucket, object, localFile, duration, maxParallel)
}
