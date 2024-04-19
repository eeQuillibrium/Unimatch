package repository

import (
	"context"

	"github.com/eeQuillibrium/Unimatch/auth_service/internal/domain/models"
	"github.com/jmoiron/sqlx"
)

const (
	defaultInt = 0
)

type AuthRepo struct {
	DB *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{DB: db}
}

func (r *AuthRepo) Register(
	ctx context.Context,
	login string,
	passHash string,
) (int, error) {

	row := r.DB.QueryRowxContext(ctx, "INSERT INTO Users (login, passhash) VALUES ($1, $2) RETURNING id", login, passHash)

	var userId int
	if err := row.Scan(&userId); err != nil {
		return defaultInt, err
	}

	return userId, nil
}
func (r *AuthRepo) Login(
	ctx context.Context,
	login string,
) (*models.User, error) {
	user := models.User{}

	err := r.DB.GetContext(ctx, &user, "SELECT * FROM Users WHERE login=$1", login)

	return &user, err
}
