package authservice

import (
	"context"
	"net/http"

	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/dto"
	httpErrors "github.com/eeQuillibrium/Unimatch/pkg/http_errors"
	"github.com/labstack/echo/v4"
)

func (a *authservice) signUpHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		var user dto.User

		if err := c.Bind(&user); err != nil {
			a.log.Errorf("signUpHandler() error: %w", err)
			return httpErrors.ParseErrors(err)
		}

		userId, err := a.gRPCCLient.Register(ctx, user.Login, user.Password)
		if err != nil {
			a.log.Errorf("signUpHandler() error: %w", err)
			return httpErrors.ParseErrors(err)
		}

		return c.JSON(http.StatusCreated, &dto.UserId{UserId: userId})
	}
}
func (a *authservice) signInHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		var user dto.User

		if err := c.Bind(&user); err != nil {
			a.log.Errorf("signInHandler() error: %w", err)
			return httpErrors.ParseErrors(err)
		}

		token, err := a.gRPCCLient.Login(ctx, user.Login, user.Password)
		if err != nil {
			a.log.Errorf("signInHandler() error: %w", err)
		}

		c.Response().Header().Set("Authorization token", token)
		c.Response().Status = http.StatusOK
		return nil
	}
}
