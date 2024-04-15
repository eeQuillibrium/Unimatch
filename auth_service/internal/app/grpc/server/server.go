package grpcserver

import (
	"github.com/eeQuillibrium/Unimatch/auth_service/internal/service"
	auth_grpc "github.com/eeQuillibrium/Unimatch/proto/gen/go/auth"
	"google.golang.org/grpc"
)

type serverAPI struct {
	auth_grpc.UnimplementedAuthServer
	authService service.Auth
}

func Register(grpcServer *grpc.Server, authService service.Auth) {
	auth_grpc.RegisterAuthServer(grpcServer, &serverAPI{authService: authService})
}
