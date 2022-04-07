package service

import (
	"fmt"
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/internal/repository"
	"github.com/AlexMykhailov1/ImageAPI/models/image"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
)

// uploadService provides business logic for uploaded files
type uploadService struct {
	imageRepos repository.ImageRepos
	cfg        *config.Config
}

// NewUploadService returns a pointer to new uploadService
func NewUploadService(ur repository.ImageRepos, cfg *config.Config) *uploadService {
	return &uploadService{imageRepos: ur, cfg: cfg}
}

func (us *uploadService) UploadImage(c *gin.Context, file *multipart.FileHeader) (uuid.UUID, error) {
	// Create new uuid for the image
	id := uuid.New()

	// Set the desired image name
	const quality = "100"
	file.Filename = image.SprintImageName(id, quality, file.Filename)

	// Create image object
	img := image.NewImage(id, file.Filename)

	// Save image object in the database
	err := us.imageRepos.AddImage(c.Request.Context(), img)
	if err != nil {
		return uuid.Nil, err
	}

	// Save image in the local storage
	fmt.Println(us.cfg.Img + file.Filename)
	err = c.SaveUploadedFile(file, us.cfg.Img+file.Filename)
	if err != nil {
		return uuid.Nil, err
	}
	// TODO add image uuid in queue
	return id, nil
}
