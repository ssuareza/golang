package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()

	conn, err := kafka.DialLeader(ctx, "tcp", "localhost:9092", "my-topic", 0)
	if err != nil {
		log.Fatalf("failed to connect to Kafka: %v", err)
	}

	defer conn.Close()

	batch := conn.ReadBatch(1, 1024)
	if err != nil {
		log.Fatalf("failed to create a message batch: %v", err)
	}

	msg, err := batch.ReadMessage()
	if err != nil {
		log.Fatalf("failed to read message: %v", err)
	}

	fmt.Println("Received message:", string(msg.Value))

	err = batch.Close()
	if err != nil {
		log.Fatalf("failed to close message batch: %v", err)
	}
}
