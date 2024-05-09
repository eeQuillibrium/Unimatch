package grpcclient

import (
	"context"
	"fmt"

	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	auth_grpc "github.com/eeQuillibrium/Unimatch/proto/gen/go/auth"
	"google.golang.org/grpc"
)

const (
	errGRPC       = "grpc: "
	defaultInt    = 0
	defaultString = ""
)

type authClient struct {
	log  *logger.Logger
	ac   auth_grpc.AuthClient
	conn *grpc.ClientConn
}

func NewAuthClient(
	log *logger.Logger,
	conn *grpc.ClientConn,
) *authClient {
	ac := auth_grpc.NewAuthClient(conn)
	return &authClient{
		log:  log,
		ac:   ac,
		conn: conn,
	}
}

func (a *authClient) Register(
	ctx context.Context,
	login string,
	password string,
) (int, error) {
	resp, err := a.ac.Register(ctx, &auth_grpc.RegisterRequest{Login: login, Password: password})
	if err != nil {
		return defaultInt, fmt.Errorf("%s Register(login, password) func: %w", errGRPC, err)
	}

	return int(resp.GetUserId()), nil
}

func (a *authClient) Login(
	ctx context.Context,
	login string,
	password string,
) (string, error) {
	resp, err := a.ac.Login(ctx, &auth_grpc.LoginRequest{Login: login, Password: password})
	if err != nil {
		return defaultString, fmt.Errorf("%s Login(login, password) func: %w", errGRPC, err)
	}

	return resp.GetToken(), nil
}
func (a *authClient) IdentifyUser(
	ctx context.Context,
	token string,
) (int, error) {
	resp, err := a.ac.IdentifyUser(ctx, &auth_grpc.IdentifyRequest{Token: token})
	if err != nil {
		return defaultInt, fmt.Errorf("%s IdentifyUser(token) func: %w", errGRPC, err)
	}

	return int(resp.GetUserId()), nil
}
