package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	grpcapp "github.com/eeQuillibrium/Unimatch/auth_service/internal/app/grpc"
	"github.com/eeQuillibrium/Unimatch/auth_service/internal/config"
	"github.com/eeQuillibrium/Unimatch/auth_service/internal/repository"
	"github.com/eeQuillibrium/Unimatch/auth_service/internal/service"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		a.cfg.PostgresDB.Host,
		a.cfg.PostgresDB.Port,
		a.cfg.PostgresDB.Username,
		os.Getenv("DB_PASSWORD"),
		a.cfg.PostgresDB.DBName,
		a.cfg.PostgresDB.SSLMode,
	)
	a.log.Info(dsn)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		a.log.Fatalf("postgres db open problem: %w", err)
	}

	authRepository := repository.NewRepository(db)
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
