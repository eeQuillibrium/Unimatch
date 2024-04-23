package main

import (
	"github.com/eeQuillibrium/Unimatch/auth_service/internal/app"
	"github.com/eeQuillibrium/Unimatch/auth_service/internal/config"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
)

func main() {
	log := logger.NewLogger()

	//if err := godotenv.Load(); err != nil {
	//	log.Fatalf("godotenv.Load() error: %w", err)
	//}

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("InitConfig() error: %w", err)
	}
	app := app.NewApp(log, cfg)

	app.Run()
}
