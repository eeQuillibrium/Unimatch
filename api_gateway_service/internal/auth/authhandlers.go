package auth

import (
	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/service"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/labstack/echo/v4"
)

type authHandlers struct {
	log         *logger.Logger
	echogroup   *echo.Group
	authService service.AuthService
}

func NewAuthHandlers(
	log *logger.Logger,
	group *echo.Group,
	service service.AuthService,
) *authHandlers {
	return &authHandlers{log: log, echogroup: group, authService: service}
}
