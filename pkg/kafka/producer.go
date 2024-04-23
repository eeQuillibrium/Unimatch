package kafka

import (
	"time"

	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/segmentio/kafka-go"
)

const (
	producerReadTimeout  = 10 * time.Second
	producerWriteTimeout = 10 * time.Second
)

type Producer struct {
	log *logger.Logger
	wr  *kafka.Writer
}

func NewProducer(log *logger.Logger, brokers []string) *Producer {
	return &Producer{log: log, wr: kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Balancer:     &kafka.RoundRobin{},
		RequiredAcks: -1,
		ReadTimeout:  producerReadTimeout,
		WriteTimeout: producerWriteTimeout,
	})}
}