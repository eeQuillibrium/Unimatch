package service

import (
	"context"
	"time"

	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/config"
	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/dto"
	kafkaClient "github.com/eeQuillibrium/Unimatch/pkg/kafka"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	kafkaMessages "github.com/eeQuillibrium/Unimatch/proto/gen/go/kafka"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type profileService struct {
	log *logger.Logger
	cfg *config.Config
	pr  *kafkaClient.Producer
}

func NewProfileService(
	log *logger.Logger, 
	cfg *config.Config, 
	pr *kafkaClient.Producer,
) *profileService {
	return &profileService{log: log, cfg: cfg, pr: pr}
}

func (s *profileService) SetProfile(
	ctx context.Context,
	profile *dto.Profile,
) error {
	profileBytes, err := proto.Marshal(&kafkaMessages.Profile{
		UserID: profile.UserId,
		Name:   profile.Name,
		Age:    profile.Age,
		About:  profile.About,
		Avatar: profile.Avatar,
	})
	if err != nil {
		return err
	}

	return s.pr.SendMessage(ctx, kafka.Message{
		Topic: s.cfg.KafkaTopics.SetProfile.TopicName,
		Value: profileBytes,
		Time:  time.Now().UTC(),
	})
}
