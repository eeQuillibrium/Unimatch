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
	cl   auth_grpc.AuthClient
	conn *grpc.ClientConn
}

func NewAuthClient(
	log *logger.Logger,
	conn *grpc.ClientConn,
) *authClient {
	cl := auth_grpc.NewAuthClient(conn)
	return &authClient{
		log:  log,
		cl:   cl,
		conn: conn,
	}
}

func (a *authClient) Register(
	ctx context.Context,
	login string,
	password string,
) (int, error) {
	resp, err := a.cl.Register(ctx, &auth_grpc.RegisterRequest{Login: login, Password: password})
	if err != nil {
		a.log.Errorf("error in grpc Register(login, password) func: %w", err)
		return defaultInt, fmt.Errorf("%s error in grpc Register(login, password) func: %w", errGRPC, err)
	}

	return int(resp.GetUserId()), nil
}

func (a *authClient) Login(
	ctx context.Context,
	login string,
	password string,
) (string, error) {
	resp, err := a.cl.Login(ctx, &auth_grpc.LoginRequest{Login: login, Password: password})
	if err != nil {
		a.log.Errorf("error in grpc Login(login, password) func: %w", err)
		return defaultString, fmt.Errorf("%s error in grpc Login(login, password) func: %w", errGRPC, err)
	}

	return resp.GetToken(), nil
}
func (a *authClient) IdentifyUser(
	ctx context.Context,
	token string,
) (int, error) {
	resp, err := a.cl.IdentifyUser(ctx, &auth_grpc.IdentifyRequest{Token: token})
	if err != nil {
		a.log.Errorf("error in grpc IdentifyUser(token) func: %w", err)
		return defaultInt, fmt.Errorf("%s error in grpc IdentifyUser(token) func: %w", errGRPC, err)
	}

	return int(resp.GetUserId()), nil
}
