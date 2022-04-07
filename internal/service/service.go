package service

import (
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
)

// UploadService stores methods of uploadService struct
type UploadService interface {
	UploadImage(c *gin.Context, file *multipart.FileHeader) (uuid.UUID, error)
}

// Services stores all service interfaces
type Services struct {
	UploadService UploadService
}

// NewServices returns a pointer to new Services
func NewServices(repos *repository.Repositories, cfg *config.Config) *Services {
	return &Services{UploadService: NewUploadService(repos.ImageRepos, cfg)}
}
