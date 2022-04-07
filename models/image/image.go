package image

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// Image name should be of format uuid.quality.imageType.7
// Quality can be 100/75/50/25. Database should store images with quality 100 only
type Image struct {
	Id   uuid.UUID
	Name string
}

// NewImage returns a pointer to new Image
func NewImage(id uuid.UUID, name string) *Image {
	return &Image{id, name}
}

// SprintImageName returns an image name in format uuid.quality.imageType ;
// imgName field should be full name of the image with its type
func SprintImageName(uuid uuid.UUID, quality, imgName string) (name string) {
	// Get image type from old name
	s := strings.Split(imgName, ".")
	imgType := s[len(s)-1]

	// Create new name
	name = fmt.Sprintf("%v.%s.%s", uuid, quality, imgType)
	return
}
