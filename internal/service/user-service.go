package service

import (
	"context"
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/internal/models/image"
	"github.com/AlexMykhailov1/ImageAPI/internal/repository"
	"github.com/google/uuid"
	"strconv"
)

// userService provides logic for all user interactions
type userService struct {
	imageRepos repository.ImageRepos
	cfg        *config.Config
}

// NewUserService returns a pointer to new userService
func NewUserService(ur repository.ImageRepos, cfg *config.Config) *userService {
	return &userService{ur, cfg}
}

// DownloadImage calls database layer to take an image name and then returns it
func (us *userService) DownloadImage(ctx context.Context, id uuid.UUID, needQuality string) (string, error) {
	// Get image object from the database
	img, err := us.imageRepos.GetImage(ctx, id)
	if err != nil {
		return "", err
	}

	// Convert needQuality to int
	nqint, err := strconv.Atoi(needQuality)
	if err != nil {
		return "", err
	}

	// Determine filename
	fileName := img.Name
	if nqint != 100 {
		fileName = image.SetNameQuality(fileName, nqint)
	}

	return fileName, nil
}
