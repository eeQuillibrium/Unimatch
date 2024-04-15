package service

import (
	"context"

	"github.com/eeQuillibrium/Unimatch/auth_service/internal/repository"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(
	ctx context.Context,
	login string,
	password string,
) (userId int, err error) {
	return 0, nil
}
func (s *AuthService) Login(
	ctx context.Context,
	login string,
	password string,
) (token string, err error) {
	return "", nil
}
func (s *AuthService) UserIdentify(
	ctx context.Context,
	token string,
) (userId int, err error) {
	return 0, nil
}
