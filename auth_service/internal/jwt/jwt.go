package jwt

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	secretKey  = "sdfslk@#$d{{}sdflvFFGSsFFFfkp[ASDASfmf]231r#@{}FSD>fs"
	defaultInt = 0
)

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"userid,omitempty"`
}

func GenerateToken(
	ctx context.Context,
	userId int,
	duration time.Duration,
) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		userId,
	})

	return token.SignedString([]byte(secretKey))
}

func ParseToken(
	ctx context.Context,
	tokenStr string,
) (int, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return secretKey, nil
		})
	if err != nil {
		return defaultInt, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return defaultInt, errors.New("token claims doesn't have *TokenClaims type")
	}

	return claims.UserId, nil
}
