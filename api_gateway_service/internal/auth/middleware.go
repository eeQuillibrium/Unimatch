package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *authHandlers) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization token")

		userId, err := a.authService.IdentifyUser(context.Background(), token)
		if err != nil {
			c.Response().WriteHeader(http.StatusUnauthorized)
			return c.JSON(http.StatusUnauthorized, token)
		}

		c.Request().Header.Set("userid", fmt.Sprintf("%d", userId))

		return next(c)
	}
}
