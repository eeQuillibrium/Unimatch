package app

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/auth"
	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/config"
	grpcapp "github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/grpc"
	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/profile"
	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/service"
	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/templates"
	"github.com/eeQuillibrium/Unimatch/pkg/kafka"
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

	a.echo.Renderer = templates.NewTemplate(a.cfg.AssetsPath)
	
	grpcApp := grpcapp.NewGRPCApp(a.log, a.cfg.GRPC.AuthPort)

	if err := grpcApp.Run(); err != nil {
		a.log.Errorf("grpcApp.Run(): %w", err)
	}

	pr := kafka.NewProducer(a.log, a.cfg.Kafka.Brokers)
	services := service.NewService(a.log, a.cfg, grpcApp.Auth, pr)

	authHandlers := auth.NewAuthHandlers(a.log, a.echo.Group("/auth"), services.Auth)
	authHandlers.MapRoutes()

	profileHandlers := profile.NewProfileHandlers(a.log, a.echo.Group("/profile", authHandlers.AuthMiddleware), services.Profile)
	profileHandlers.MapRoutes()

	go func() {
		if err := a.runHttpServer(); err != nil {
			a.log.Errorf("error with runHttoServer %w", err)
			cancel()
		}
	}()

	<-ctx.Done()

	if err := a.echo.Shutdown(ctx); err != nil {
		a.log.Warn("echo shutdown: %w", err)
	}
	
	return nil
}
