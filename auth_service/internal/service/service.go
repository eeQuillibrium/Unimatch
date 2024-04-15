package service

import (
	"context"

	"github.com/eeQuillibrium/Unimatch/auth_service/internal/repository"
)

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
	UserIdentify(
		ctx context.Context,
		token string,
	) (userId int, err error)
}

type Service struct {
	AuthService Auth
}


func NewService(authRepos *repository.Repository) *Service {
	return &Service{AuthService: NewAuthService(authRepos.AuthRepository)}
}