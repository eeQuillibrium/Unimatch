package authservice

import (
	grpcapp "github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/grpc"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/labstack/echo/v4"
)

type authservice struct {
	log        *logger.Logger
	echogroup  *echo.Group
	gRPCCLient grpcapp.AuthGRPC
}

func NewAuthService(
	log *logger.Logger,
	echogroup *echo.Group,
	gRPCCLient grpcapp.AuthGRPC,
) *authservice {
	return &authservice{
		log: log,
		echogroup: echogroup,
		gRPCCLient: gRPCCLient,
	}
}

