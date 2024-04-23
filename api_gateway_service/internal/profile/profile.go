package profile

import (
	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/service"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/labstack/echo/v4"
)

type profileHandlers struct {
	log   *logger.Logger
	group *echo.Group
	profileService service.ProfileService
}

func NewProfileHandlers(
	log *logger.Logger,
	group *echo.Group,
	profileService service.ProfileService,
) *profileHandlers {
	return &profileHandlers{log: log, group: group, profileService: profileService}
}
