package repository

import (
	"context"
	"database/sql"
	"github.com/AlexMykhailov1/ImageAPI/internal/models/image"
	"github.com/google/uuid"
)

// ImageRepos stores all method of imageRepos struct
type ImageRepos interface {
	AddImage(ctx context.Context, img *image.Image) error
	GetImage(ctx context.Context, id uuid.UUID) (*image.Image, error)
}

// Repositories stores all repository interfaces
type Repositories struct {
	ImageRepos ImageRepos
}

// NewRepositories returns a pointer to new Repositories
func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{ImageRepos: NewImageRepos(db)}
}
