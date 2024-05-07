package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func NewKafkaConn(
	ctx context.Context,
	cfg *Config,
) (*kafka.Conn, error) {
	return kafka.DialContext(ctx, "tcp", cfg.Brokers[0])
}
