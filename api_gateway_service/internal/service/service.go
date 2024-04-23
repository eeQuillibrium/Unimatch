package service

import (
	"context"

	grpcapp "github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/grpc"
	"github.com/eeQuillibrium/Unimatch/pkg/kafka"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
)

type AuthService interface {
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
	IdentifyUser(
		ctx context.Context,
		token string,
	) (userId int, err error)
}
type ProfileService interface {
	
}

type Service struct {
	Auth    AuthService
	Profile ProfileService
}

func NewService(
	log *logger.Logger,
	gRPCClient grpcapp.AuthGRPC,
	pr *kafka.Producer,
) *Service {
	return &Service{
		Auth:    NewAuthService(log, gRPCClient),
		Profile: NewProfileService(log, pr),
	}
}
