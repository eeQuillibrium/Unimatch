package profile

import (
	"context"
	"net/http"
	"strconv"

	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/dto"
	httpErrors "github.com/eeQuillibrium/Unimatch/pkg/http_errors"

	"github.com/labstack/echo/v4"
)

func (h *profileHandlers) setProfileHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		mForm, err := c.MultipartForm()
		if err != nil {
			return httpErrors.ParseErrors(err)
		}

		profile, err := dto.AccessProfile(ctx, mForm)
		if err != nil {
			return httpErrors.ParseErrors(err)
		}
		userID, err := strconv.Atoi(c.Request().Header.Get("userid"))
		if err != nil {
			return httpErrors.ParseErrors(err)
		}
		profile.UserId = int64(userID)

		if err := h.profileService.SetProfile(ctx, profile); err != nil {
			httpErrors.ParseErrors(err)
		}

		return nil
	}
}

func (h *profileHandlers) getProfileHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h *profileHandlers) profileHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Render(http.StatusOK, "profileForm.html", nil); err != nil {
			return httpErrors.ParseErrors(err)
		}

		return nil
	}
}
