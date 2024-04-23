package app

import (
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/config"
)

type app struct {
	log *logger.Logger
	cfg *config.Config
}

func New(log *logger.Logger, cfg *config.Config) *app {
	return &app{log: log, cfg: cfg}
}
