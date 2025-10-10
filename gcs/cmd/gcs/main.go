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

	// Download the object
	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to open GCS object: %w", err))
	}
	defer rc.Close()

	// Create local file
	f, err := os.Create(localFile)
	if err != nil {
		panic(fmt.Errorf("failed to create local file: %w", err))
	}
	defer f.Close()

	// Measure download time
	start := time.Now()
	if _, err := io.Copy(f, rc); err != nil {
		panic(fmt.Errorf("failed to copy data: %w", err))
	}
	duration := time.Since(start)

	// Success message
	fmt.Printf("Downloaded gs://%s/%s to %s in %s\n", bucket, object, localFile, duration)
}
