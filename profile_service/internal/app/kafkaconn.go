package app

import (
	"context"

	kafkaClient "github.com/eeQuillibrium/Unimatch/pkg/kafka"
)

func (a *App) setKafkaConn(
	ctx context.Context,
	cfg *kafkaClient.Config,
) error {
	kafkaConn, err := kafkaClient.NewKafkaConn(ctx, cfg)
	if err != nil {
		return err
	}
	a.kafkaConn = kafkaConn

	brokers, err := kafkaConn.Brokers()
	if err != nil {
		return err
	}

	a.log.Infof("connected to brokers: %+v", brokers)

	return nil
}

func (a *App) getKafkaTopics() []string {
	return []string{
		a.cfg.KafkaTopics.SetProfile.TopicName,
	}
}
