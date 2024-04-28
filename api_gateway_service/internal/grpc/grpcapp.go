package grpcapp

import (
	"context"
	"fmt"

	grpcclient "github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/grpc/client"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthGRPC interface {
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

type grpcapp struct {
	log      *logger.Logger
	authPort int
	Auth     AuthGRPC
}

func NewGRPCApp(
	log *logger.Logger,
	authPort int,
) *grpcapp {
	return &grpcapp{
		log:      log,
		authPort: authPort,
	}
}

func (a *grpcapp) Run() error {
	conn, err := grpc.Dial(fmt.Sprintf("auth_service:%d", a.authPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	a.Auth = grpcclient.NewAuthClient(a.log, conn)
	return nil
}
