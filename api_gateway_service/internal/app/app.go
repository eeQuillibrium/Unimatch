package app

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/config"
	grpcapp "github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/grpc"
	authservice "github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/services/auth"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/labstack/echo/v4"
)

type app struct {
	echo *echo.Echo
	cfg  *config.Config
	log  *logger.Logger
}

func NewApp(
	log *logger.Logger,
	cfg *config.Config,
) *app {
	return &app{echo: echo.New(), log: log, cfg: cfg}
}

func (a *app) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	grpcApp := grpcapp.NewGRPCApp(a.log, a.cfg.GRPC.AuthPort)

	if err := grpcApp.Run(); err != nil {
		a.log.Errorf("error with runHttoServer %w", err)
	}

	authService := authservice.NewAuthService(a.log, a.echo.Group("/auth"), grpcApp.Auth)
	authService.MapRoutes()

	go func() {
		if err := a.runHttpServer(); err != nil {
			a.log.Errorf("error with runHttoServer %w", err)
			cancel()
		}
	}()

	<-ctx.Done()

	return nil
}
