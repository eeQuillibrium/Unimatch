package service

import (
	"context"
	"errors"
	"time"

	"github.com/eeQuillibrium/Unimatch/auth_service/internal/jwt"
	"github.com/eeQuillibrium/Unimatch/auth_service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

const (
	emptyPass  = ""
	emptyLogin = ""

	defaultInt    = 0
	defaultString = ""

	salt = "GFH#$!@#rglksdpSDFSDlmvk4r5l3[]xvcvx"
	cost = 3
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
	if password == emptyPass {
		return defaultInt, errors.New("Register() empty password")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return defaultInt, errors.New("Register() generatePass error")
	}

	return s.repo.Register(ctx, login, string(passHash))
}
func (s *AuthService) Login(
	ctx context.Context,
	login string,
	password string,
) (string, error) {
	if password == emptyPass {
		return defaultString, errors.New("Register() empty password")
	}

	user, err := s.repo.Login(ctx, login)
	if err != nil {
		return defaultString, errors.New("Login() repo problem")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(password)); err != nil {
		return defaultString, errors.New("Register() generatePass error")
	}

	return jwt.GenerateToken(ctx, user.Id, time.Hour)
}
func (s *AuthService) UserIdentify(
	ctx context.Context,
	token string,
) (int, error) {
	return jwt.ParseToken(ctx, token)
}
