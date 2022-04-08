package optimizer

import (
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/internal/models/image"
	"github.com/h2non/bimg"
)

// SaveLessQuality reads image with given name and saves 3 same images with quality 25%/50%/75% using SetQuality function
func SaveLessQuality(cfg *config.Config, name string) error {
	// Read image into the buffer
	origImg, err := bimg.Read(cfg.Path.Img + name)
	if err != nil {
		return err
	}
	// Process and save 3 new images from original image
	qualities := [3]int{75, 50, 25}
	for _, quality := range qualities {
		// Process new image
		newImg, err := processQuality(origImg, quality)
		if err != nil {
			return err
		}
		// Change name for the new image
		name = image.SetNameQuality(name, quality)
		// Save new image
		err = bimg.Write(cfg.Path.Img+name, newImg)
		if err != nil {
			return err
		}
	}
	return nil
}

// processQuality processes image in the buffer and returns the same image with the given quality
func processQuality(buffer []byte, quality int) ([]byte, error) {
	// Processing options
	options := bimg.Options{
		Quality: quality,
	}
	// Process image
	newImage, err := bimg.NewImage(buffer).Process(options)
	if err != nil {
		return nil, err
	}
	return newImage, nil
}
