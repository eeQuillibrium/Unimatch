package kafka

import (
	"context"
	"sync"

	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type Worker func(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)

type ConsumerGroup interface {
	ConsumeTopics(
		ctx context.Context,
		topics []string,
		poolSize int,
		worker Worker,
	)
}

type consumerGroup struct {
	log     *logger.Logger
	brokers []string
	groupID string
}

func (c *consumerGroup) NewReader(brokers []string, groupTopics []string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:                brokers,
		GroupID:                groupID,
		GroupTopics:            groupTopics,
		MinBytes:               minBytes,
		MaxBytes:               maxBytes,
		QueueCapacity:          queueCapacity,
		HeartbeatInterval:      heartbeatInterval,
		CommitInterval:         commitInterval,
		PartitionWatchInterval: partitionWatchInterval,
		MaxAttempts:            maxAttempts,
		MaxWait:                maxWait,
		Dialer: &kafka.Dialer{
			Timeout: dialTimeout,
		},
	})
}

func NewConsumerGroup(
	log *logger.Logger,
	brokers []string,
	groupID string,
) *consumerGroup {
	return &consumerGroup{log: log, brokers: brokers, groupID: groupID}
}

func (c *consumerGroup) ConsumeTopics(
	ctx context.Context,
	topics []string,
	poolSize int,
	worker Worker,
) {
	consumer := c.NewReader(c.brokers, topics, c.groupID)

	defer func() {
		if err := consumer.Close(); err != nil {
			c.log.Warnf("consumer.Close() %w", err)
		}
	}()
	c.log.Infof("run consumerGroup with topic: %s", topics)

	wg := &sync.WaitGroup{}
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go worker(ctx, consumer, wg, i)
	}

	wg.Wait()
}
