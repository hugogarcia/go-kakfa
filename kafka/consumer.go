package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context, topic string) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{KafkaServerAddress},
		Topic:     topic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	fmt.Println("Consumer started!")

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	fmt.Println("Consumer finished")

	return nil
}
