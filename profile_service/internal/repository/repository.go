package repository

import (
	"context"

	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/config"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/domain/models"
	"github.com/jmoiron/sqlx"
)

type ProfileProvider interface {
	StoreProfile(
		ctx context.Context,
		profile *models.Profile,
	) error
}

type Repository struct {
	ProfileProvider
}

func NewRepository(
	log *logger.Logger,
	cfg *config.Config,
	db *sqlx.DB,
) *Repository {
	return &Repository{
		ProfileProvider: NewProfileProvider(log, cfg, db),
	}
}
