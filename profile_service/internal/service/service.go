package service

import (
	"context"

	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/repository"
	kafkaMessages "github.com/eeQuillibrium/Unimatch/proto/gen/go/kafka"
)

type ProfileProvider interface {
	StoreProfile(
		ctx context.Context,
		profile *kafkaMessages.Profile,
	) error
}
type Service struct {
	ProfileProvider
}

func NewService(
	log *logger.Logger,
	repo *repository.Repository,
) *Service {
	return &Service{
		ProfileProvider: NewProfileProvider(log, repo.ProfileProvider),
	}
}
