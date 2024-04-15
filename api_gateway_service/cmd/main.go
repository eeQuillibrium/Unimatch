package main

import (
	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/app"
	"github.com/eeQuillibrium/Unimatch/api_gateway_service/internal/config"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {

	log := logger.NewLogger()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("env params loading problem: %v", err)
	}

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	app := app.NewApp(log, cfg)

	app.Run()
}
