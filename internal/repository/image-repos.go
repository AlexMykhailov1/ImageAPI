package repository

import (
	"context"
	"database/sql"
	"github.com/AlexMykhailov1/ImageAPI/internal/models/image"
	"log"
)

// imageRepos provides database communication logic for images
type imageRepos struct {
	db *sql.DB
}

// NewImageRepos returns a pointer to new imageRepos
func NewImageRepos(db *sql.DB) *imageRepos {
	return &imageRepos{db: db}
}

// AddImage adds new image object to the database
func (ir *imageRepos) AddImage(ctx context.Context, img *image.Image) error {
	query := `INSERT INTO images(id, name) VALUES ($1,$2);`
	result, err := ir.db.ExecContext(ctx, query, img.Id, img.Name)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}
