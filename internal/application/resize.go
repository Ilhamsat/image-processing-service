package application

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image-processing/internal/domain"
	"os"
)

//func ResizeImage(inputPath, outputPath string, width, height int) error {
//	// Open the input image file
//	inputFile, err := os.Open(inputPath)
//	if err != nil {
//		return fmt.Errorf("failed to open input file: %v", err)
//	}
//	defer inputFile.Close()
//
//	// Decode the input image
//	inputImage, err := imaging.Decode(inputFile)
//	if err != nil {
//		return fmt.Errorf("failed to decode input image: %v", err)
//	}
//
//	// Resize the image
//	resizedImage := imaging.Resize(inputImage, width, height, imaging.Lanczos)
//
//	// Save the resized image
//	err = imaging.Save(resizedImage, outputPath)
//	if err != nil {
//		return fmt.Errorf("failed to save output image: %v", err)
//	}
//
//	return nil
//}

func ResizeImage(inputPath, outputPath string, width, height int) error {
	// Open the input image file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer inputFile.Close()

	// Decode the input image
	inputImage, _, err := image.Decode(inputFile)
	if err != nil {
		return fmt.Errorf("failed to decode input image: %v", err)
	}

	// Create a new Image entity
	imageEntity := domain.NewImage("image-id", "jpeg", nil, inputImage, nil)

	// Perform image resizing
	resizedImage := imaging.Resize(inputImage, width, height, imaging.Lanczos)

	// Update the Image entity with resized data
	imageEntity.Width = width
	imageEntity.Height = height
	imageEntity.Data = nil // Set to nil since we don't store the image data in memory

	// Save the resized image
	err = imaging.Save(resizedImage, outputPath)
	if err != nil {
		return fmt.Errorf("failed to save output image: %v", err)
	}

	return nil
}
