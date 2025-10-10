package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

func main() {
	ctx := context.Background()

	bucketName := "file-321"  // set your bucket name
	objectName := "file.bin"  // set your object name
	localFile := "./file.bin" // set your local file path

	client, err := storage.NewGRPCClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Minute*10)
	defer cancel()

	rc, err := client.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		log.Fatalf("Failed to open GCS object: %v", err)
	}
	defer rc.Close()

	f, err := os.Create(localFile)
	if err != nil {
		log.Fatalf("Failed to create local file: %v", err)
	}
	defer f.Close()

	start := time.Now()
	if _, err := io.Copy(f, rc); err != nil {
		log.Fatalf("Failed to copy data: %v", err)
	}
	duration := time.Since(start)

	fmt.Printf("Downloaded gs://%s/%s to %s in %s\n", bucketName, objectName, localFile, duration)
}
