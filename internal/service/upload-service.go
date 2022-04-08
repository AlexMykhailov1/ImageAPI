package service

import (
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/internal/models/image"
	"github.com/AlexMykhailov1/ImageAPI/internal/rabbit"
	"github.com/AlexMykhailov1/ImageAPI/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
)

// uploadService provides business logic for uploaded files
type uploadService struct {
	imageRepos repository.ImageRepos
	cfg        *config.Config
	rb         *rabbit.Rabbit
}

// NewUploadService returns a pointer to new uploadService
func NewUploadService(ur repository.ImageRepos, cfg *config.Config, rb *rabbit.Rabbit) *uploadService {
	return &uploadService{imageRepos: ur, cfg: cfg, rb: rb}
}

// UploadImage implements all the logic after the image was uploaded to the server
func (us *uploadService) UploadImage(c *gin.Context, file *multipart.FileHeader) (uuid.UUID, error) {
	// Create new uuid for the image
	id := uuid.New()

	// Set the desired image name
	const quality = 100
	file.Filename = image.SprintImageName(id, quality, file.Filename)

	// Create image object
	img := image.NewImage(id, file.Filename)

	// Call database repository layer to save image object in the database
	if err := us.imageRepos.AddImage(c.Request.Context(), img); err != nil {
		return uuid.Nil, err
	}

	// Save image in the local storage
	if err := c.SaveUploadedFile(file, us.cfg.Path.Img+file.Filename); err != nil {
		return uuid.Nil, err
	}

	// Send image id to the queue
	if err := us.rb.SendImgID(us.cfg, img.Name); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
