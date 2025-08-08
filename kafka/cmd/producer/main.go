package main

import (
	"context"
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

	_, err = conn.WriteMessages(kafka.Message{
		Value: []byte("hello world"),
	})
	if err != nil {
		log.Fatalf("failed to write messages: %v", err)
	}
}
