package grpcserver

import (
	"context"

	"github.com/eeQuillibrium/Unimatch/profile_service/internal/service"
	profile_grpc "github.com/eeQuillibrium/Unimatch/proto/gen/go/profile"
	"google.golang.org/grpc"
)

type serverAPI struct {
	service service.ProfileProvider
	profile_grpc.UnimplementedProfileServer
}

func Register(
	gRPCServ *grpc.Server,
	service service.ProfileProvider,
) {
	profile_grpc.RegisterProfileServer(gRPCServ, &serverAPI{service: service})
}

func (s *serverAPI) GetProfile(
	ctx context.Context,
	req *profile_grpc.GetProfileReq,
) (*profile_grpc.GetProfileResp, error) {
	return nil, nil
}
