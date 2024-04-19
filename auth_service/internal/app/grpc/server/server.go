package grpcserver

import (
	"context"
	"log"

	"github.com/eeQuillibrium/Unimatch/auth_service/internal/service"
	auth_grpc "github.com/eeQuillibrium/Unimatch/proto/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	auth_grpc.UnimplementedAuthServer
	authService service.Auth
}

func Register(grpcServer *grpc.Server, authService service.Auth) {
	auth_grpc.RegisterAuthServer(grpcServer, &serverAPI{authService: authService})
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *auth_grpc.RegisterRequest,
) (*auth_grpc.RegisterResponse, error) {
	userId, err := s.authService.Register(ctx, req.GetLogin(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "register error: %v", err)
	}

	return &auth_grpc.RegisterResponse{UserId: int64(userId)}, nil
}
func (s *serverAPI) Login(
	ctx context.Context,
	req *auth_grpc.LoginRequest,
) (*auth_grpc.LoginResponse, error) {
	token, err := s.authService.Login(ctx, req.GetLogin(), req.GetPassword())
	if err != nil {
		log.Printf("token err: %v", err)
		return nil, status.Errorf(codes.Internal, "login error: %v", err)
	}
	return &auth_grpc.LoginResponse{Token: token}, nil
}
func (s *serverAPI) IdentifyUser(
	ctx context.Context,
	req *auth_grpc.IdentifyRequest,
) (*auth_grpc.IdentifyResponse, error) {
	userId, err := s.authService.UserIdentify(ctx, req.GetToken())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "identify error: %v", err)
	}
	return &auth_grpc.IdentifyResponse{UserId: int64(userId)}, nil
}
