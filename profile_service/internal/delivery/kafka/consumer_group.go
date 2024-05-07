package kafka

import (
	"context"
	"sync"

	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/config"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/service"
	"github.com/segmentio/kafka-go"
)

const (
	PoolSize = 12
)

type messageReader struct {
	log      *logger.Logger
	cfg      *config.KafkaTopics
	services *service.Service
}

func NewMessageReader(log *logger.Logger, cfg *config.KafkaTopics, services *service.Service) *messageReader {
	return &messageReader{
		log:      log,
		cfg:      cfg,
		services: services,
	}
}

func (w *messageReader) MessageReaderWorker(
	ctx context.Context,
	r *kafka.Reader,
	wg *sync.WaitGroup,
	workerID int,
) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			{
				return
			}
		default:
		}
		m, err := r.FetchMessage(ctx)
		if err != nil {
			w.log.Warnf("workerID: %d, err: %v", workerID, err)
			continue
		}

		switch m.Topic {
		case w.cfg.SetProfile.TopicName:
			w.SetProfileProcessor(ctx, r, m)
		}
	}
}
