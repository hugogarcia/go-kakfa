package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

const (
	UsersTopic         = "users"
	KafkaServerAddress = "localhost:29092"
	partition          = 0
)

func SetupProducer(ctx context.Context, topic string) (*kafka.Conn, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", KafkaServerAddress, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	//conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	return conn, nil
}
