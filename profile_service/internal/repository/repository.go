package repository

import (
	"context"

	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/config"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/domain/models"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type ProfileProvider interface {
	StoreProfile(
		ctx context.Context,
		profile *models.Profile,
	) error
	GetProfile(
		ctx context.Context,
		userID int,
	) (*models.Profile, error)
}

type Repository struct {
	ProfileProvider
}

func NewRepository(
	log *logger.Logger,
	cfg *config.Config,
	db *sqlx.DB,
	rdb *redis.Client,
) *Repository {
	return &Repository{
		ProfileProvider: NewProfileProvider(log, cfg, db, rdb),
	}
}
