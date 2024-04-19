package repository

import (
	"context"

	"github.com/eeQuillibrium/Unimatch/auth_service/internal/domain/models"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	Register(
		ctx context.Context,
		login string,
		passHash string,
	) (userId int, err error)
	Login(
		ctx context.Context,
		login string,
	) (user *models.User, err error)
}

type Repository struct {
	AuthRepository Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{AuthRepository: NewAuthRepository(db)}
}
