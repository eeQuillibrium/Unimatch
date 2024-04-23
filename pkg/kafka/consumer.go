package kafka

import (
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type consumer struct {
	log *logger.Logger
	r   *kafka.Reader
}

func NewConsumer(log *logger.Logger, brokers []string, groupID, topic string) *consumer {
	return &consumer{log: log, r: kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: groupID,
		Topic:   topic,
	})}
}
