package repository

import (
	"context"
	"database/sql"

	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/config"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/domain/models"
	"github.com/jmoiron/sqlx"
)

type profileProvider struct {
	log *logger.Logger
	cfg *config.Config
	db  *sqlx.DB
}

func NewProfileProvider(
	log *logger.Logger,
	cfg *config.Config,
	db *sqlx.DB,
) *profileProvider {
	return &profileProvider{log: log, cfg: cfg, db: db}
}

func (r *profileProvider) StoreProfile(
	ctx context.Context,
	profile *models.Profile,
) error {
	tx := r.db.MustBeginTx(ctx, &sql.TxOptions{})

	tx.MustExecContext(ctx,
		"INSERT INTO Profiles (id, name, user_age, about, img_path)"+
			"VALUES ($1, $2, $3, $4, $5)",
		profile.UserID, profile.Name, profile.Age, profile.About, profile.ImagePath)

	return tx.Commit()
}
