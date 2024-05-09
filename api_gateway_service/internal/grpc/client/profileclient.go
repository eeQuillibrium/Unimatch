package grpcclient

import (
	"context"

	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/dto"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	profile_grpc "github.com/eeQuillibrium/Unimatch/proto/gen/go/profile"
	"google.golang.org/grpc"
)

type profileClient struct {
	log  *logger.Logger
	pc   profile_grpc.ProfileClient
	conn *grpc.ClientConn
}

func NewProfileClient(
	log *logger.Logger,
	conn *grpc.ClientConn,
) *profileClient {
	pc := profile_grpc.NewProfileClient(conn)
	return &profileClient{
		log:  log,
		pc:   pc,
		conn: conn,
	}
}

func (c *profileClient) GetProfile(
	ctx context.Context,
	userID int,
) (*dto.GetProfile, error) {
	profileGRPC, err := c.pc.GetProfile(ctx, &profile_grpc.GetProfileReq{UserID: int64(userID)})
	if err != nil {
		return nil, err
	}

	return &dto.GetProfile{
		Name:    profileGRPC.GetName(),
		Age:     profileGRPC.GetAge(),
		About:   profileGRPC.GetAbout(),
		ImgPath: profileGRPC.GetImgPath(),
	}, nil
}
