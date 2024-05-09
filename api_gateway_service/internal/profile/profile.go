package profile

import (
	grpcapp "github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/grpc"
	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/service"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/labstack/echo/v4"
)

type profileHandlers struct {
	log            *logger.Logger
	group          *echo.Group
	profileService service.ProfileService
	gRPCApp        grpcapp.ProfileGRPC
}

func NewProfileHandlers(
	log *logger.Logger,
	group *echo.Group,
	profileService service.ProfileService,
) *profileHandlers {
	return &profileHandlers{log: log, group: group, profileService: profileService}
}
