package main

import (
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/app"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	log := logger.NewLogger()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("godotenv.Load() %w", err)
	}

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("config.InitConfig() %w", err)
	}

	appl := app.NewApp(log, cfg)

	if err := appl.Run(); err != nil {
		log.Fatalf("app.Run(): %v", err)
	}
}
