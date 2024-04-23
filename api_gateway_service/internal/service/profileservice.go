package service

import (
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/eeQuillibrium/Unimatch/pkg/kafka"
)

type profileService struct {
	log *logger.Logger
	pr  *kafka.Producer
}

func NewProfileService(log *logger.Logger, pr *kafka.Producer) *profileService {
	return &profileService{log: log, pr: pr}
}


