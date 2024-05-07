package service

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/domain/models"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/repository"
	kafkaMessages "github.com/eeQuillibrium/Unimatch/proto/gen/go/kafka"
)

type profileProvider struct {
	log  *logger.Logger
	repo repository.ProfileProvider
}

func NewProfileProvider(
	log *logger.Logger,
	repo repository.ProfileProvider,
) *profileProvider {
	return &profileProvider{log: log, repo: repo}
}

func (s *profileProvider) StoreProfile(
	ctx context.Context,
	profile *kafkaMessages.Profile,
) error {
	imgPath := fmt.Sprintf("images/av%05d.jpg", profile.UserID)

	file, err := os.Create(imgPath)
	if err != nil {
		s.log.Warnf("os.Create(): %v", err)
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(profile.Avatar))
	if err != nil {
		s.log.Warnf("image.Decode(): %v", err)
		return err
	}

	var opts jpeg.Options
	opts.Quality = 100

	if err := jpeg.Encode(file, img, &opts); err != nil {
		s.log.Warnf("jpeg.Encode(): %v", err)
		return err
	}

	profileModel := models.AccessProfile(profile, imgPath)

	if err := s.repo.StoreProfile(ctx, profileModel); err != nil {
		s.log.Warnf("repo.StoreProfile(): %v", err)
	}
	
	return nil
}
