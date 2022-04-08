package image

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strconv"
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

// SprintImageName returns an image name in format uuid.quality.imageType.
// imgName field is the initial image name
func SprintImageName(uuid uuid.UUID, quality int, imgName string) (name string) {
	// Get image type from old name
	s := strings.Split(imgName, ".")
	imgType := s[len(s)-1]

	// Create new name
	name = fmt.Sprintf("%v.%v.%s", uuid, quality, imgType)
	return
}

// GetQuality returns image quality from image name in format uuid.quality.imageType
func GetQuality(name string) (*int, error) {
	s := strings.Split(name, ".")
	q := s[len(s)-2]
	// Check for error
	if q != "100" && q != "75" && q != "50" && q != "25" {
		err := errors.New("failed to retrieve quality")
		return nil, err
	}
	// Convert string to int
	quality, err := strconv.Atoi(q)
	if err != nil {
		return nil, err
	}
	return &quality, nil
}

// SetNameQuality changes quality in given name to given quality
func SetNameQuality(name string, newQuality int) string {
	s := strings.Split(name, ".")
	name = fmt.Sprintf("%v.%v.%v", s[0], string(rune(newQuality)), s[len(s)-1])
	return name
}
