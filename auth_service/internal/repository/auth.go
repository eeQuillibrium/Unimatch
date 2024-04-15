package repository

import "context"

type AuthRepo struct {
}

func NewAuthRepository() *AuthRepo {
	return &AuthRepo{}
}

func (r *AuthRepo) Register(
	ctx context.Context,
	login string,
	password string,
) (userId int, err error) {
	return 0, nil
}
func (r *AuthRepo) Login(
	ctx context.Context,
	login string,
	password string,
) (token string, err error) {
	return "", nil
}
