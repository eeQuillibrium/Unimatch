package service

import (
	"context"

	grpcapp "github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/grpc"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
)

type authService struct {
	log        *logger.Logger
	gRPCCLient grpcapp.AuthGRPC
}

func NewAuthService(
	log *logger.Logger,
	gRPCCLient grpcapp.AuthGRPC,
) *authService {
	return &authService{
		log:        log,
		gRPCCLient: gRPCCLient,
	}
}
func (s *authService) Register(
	ctx context.Context,
	login string,
	password string,
) (userId int, err error) {
	return s.gRPCCLient.Register(ctx, login, password)
}
func (s *authService) Login(
	ctx context.Context,
	login string,
	password string,
) (token string, err error) {
	return s.gRPCCLient.Login(ctx, login, password)
}
func (s *authService) IdentifyUser(
	ctx context.Context,
	token string,
) (userId int, err error) {
	return s.gRPCCLient.IdentifyUser(ctx, token)
}
