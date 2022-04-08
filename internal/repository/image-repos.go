package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/AlexMykhailov1/ImageAPI/internal/models/image"
	"github.com/google/uuid"
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

	// Execute query
	result, err := ir.db.ExecContext(ctx, query, img.Id, img.Name)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

// GetImage returns an image object with given uuid
func (ir *imageRepos) GetImage(ctx context.Context, id uuid.UUID) (*image.Image, error) {
	var img image.Image
	query := `SELECT * FROM images WHERE id = $1;`

	// Execute query
	row := ir.db.QueryRowContext(ctx, query, id)

	// Bind row to the image object
	err := row.Scan(&img.Id, &img.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.New("no image with such id")
		}
		return nil, err
	}

	return &img, nil
}
