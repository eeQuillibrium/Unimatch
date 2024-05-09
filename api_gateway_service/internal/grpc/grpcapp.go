package grpcapp

import (
	"context"
	"fmt"

	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/config"
	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/dto"
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

type ProfileGRPC interface {
	GetProfile(
		ctx context.Context,
		userID int,
	) (*dto.GetProfile, error)
}

type grpcapp struct {
	log     *logger.Logger
	cfg     *config.Config
	Auth    AuthGRPC
	Profile ProfileGRPC
}

func NewGRPCApp(
	log *logger.Logger,
	cfg *config.Config,
) *grpcapp {
	return &grpcapp{
		log: log,
		cfg: cfg,
	}
}

func (a *grpcapp) Run() error {
	authConn, err := grpc.Dial(fmt.Sprintf("%s:%d", a.cfg.GRPC.AuthHost, a.cfg.GRPC.AuthPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	a.Auth = grpcclient.NewAuthClient(a.log, authConn)

	profileConn, err := grpc.Dial(fmt.Sprintf("%s:%d", a.cfg.GRPC.ProfileHost, a.cfg.GRPC.ProfilePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	a.Profile = grpcclient.NewProfileClient(a.log, profileConn)
	return nil
}
