package kafka

import (
	"context"

	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	log *logger.Logger
	wr  *kafka.Writer
}

func NewProducer(
	log *logger.Logger,
	brokers []string,
) *Producer {
	return &Producer{log: log, wr: kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Balancer:     &kafka.RoundRobin{},
		RequiredAcks: producerRequiredAcks,
		ReadTimeout:  producerReadTimeout,
		WriteTimeout: producerWriteTimeout,
		MaxAttempts:  producerMaxAttempts,
	})}
}

func (p *Producer) SendMessage(
	ctx context.Context,
	msgs ...kafka.Message,
) error {
	return p.wr.WriteMessages(ctx, msgs...)
}
