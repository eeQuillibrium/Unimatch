package repository

import "context"

type Auth interface {
	Register(
		ctx context.Context,
		login string,
		password string,
	) (userId int, err error)
	Login(
		ctx context.Context,
		login string,
		password string,
	) (token string, err error)
}

type Repository struct {
	AuthRepository Auth
}

func NewRepository() *Repository {
	return &Repository{AuthRepository: NewAuthRepository()}
}
