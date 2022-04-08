package service

import (
	"context"
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/internal/rabbit"
	"github.com/AlexMykhailov1/ImageAPI/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
)

// UploadService stores methods of uploadService struct
type UploadService interface {
	UploadImage(c *gin.Context, file *multipart.FileHeader) (uuid.UUID, error)
}

// UserService stores methods of userService struct
type UserService interface {
	DownloadImage(ctx context.Context, id uuid.UUID, quality string) (string, error)
}

// Services stores all service interfaces
type Services struct {
	UploadService UploadService
	UserService   UserService
}

// NewServices returns a pointer to new Services
func NewServices(repos *repository.Repositories, cfg *config.Config, rb *rabbit.Rabbit) *Services {
	return &Services{
		UploadService: NewUploadService(repos.ImageRepos, cfg, rb),
		UserService:   NewUserService(repos.ImageRepos, cfg),
	}
}
