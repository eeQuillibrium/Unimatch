package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/config"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/domain/models"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type profileProvider struct {
	log *logger.Logger
	cfg *config.Config
	db  *sqlx.DB
	rdb *redis.Client
}

func NewProfileProvider(
	log *logger.Logger,
	cfg *config.Config,
	db *sqlx.DB,
	rdb *redis.Client,
) *profileProvider {
	return &profileProvider{log: log, cfg: cfg, db: db, rdb: rdb}
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

func (r *profileProvider) GetProfile(
	ctx context.Context,
	userID int,
) (*models.Profile, error) {
	rCmd := r.rdb.HGetAll(ctx, fmt.Sprintf("profile:%d", userID))
	if rCmd.Err() == nil { // cash hit
		return models.AccessRedisProfile(rCmd.Val(), userID)
	}

	var profile models.Profile

	if err := r.db.GetContext(ctx, &profile, "SELECT (id, name. user_age, about, img_path)"+
		"WHERE id = $1", userID); err != nil {
		return nil, err
	}

	if err := r.rdb.HSet(ctx, fmt.Sprintf("profile:%d", userID),
		"name", profile.Name,
		"age", fmt.Sprintf("%d", profile.Age),
		"about", profile.About,
		"imagepath", profile.ImagePath).Err(); err != nil {
		return &profile, fmt.Errorf("redis: %v", err)
	}

	return &profile, nil
}
