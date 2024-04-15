package app

import (
	"context"
	"os/signal"
	"syscall"

	grpcapp "github.com/eeQuillibrium/Unimatch/auth_service/internal/app/grpc"
	"github.com/eeQuillibrium/Unimatch/auth_service/internal/config"
	"github.com/eeQuillibrium/Unimatch/auth_service/internal/repository"
	"github.com/eeQuillibrium/Unimatch/auth_service/internal/service"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
)

type app struct {
	log *logger.Logger
	cfg *config.Config
}

func NewApp(
	log *logger.Logger,
	cfg *config.Config,
) *app {
	return &app{
		log: log,
		cfg: cfg,
	}
}

func (a *app) Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	
	authRepository := repository.NewRepository()
	authService := service.NewService(authRepository)
	grpcApp := grpcapp.NewApp(a.log, authService.AuthService, a.cfg.GRPC.Serverport)

	go func() {
		if err := grpcApp.Run(ctx); err != nil {
			a.log.Fatalf("Run() error: %w", err)
		}

		cancel()
	}()

	<-ctx.Done()
}
