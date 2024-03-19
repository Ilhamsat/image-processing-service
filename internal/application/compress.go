package application

import (
	"fmt"
	"image"
	"image-processing/internal/domain"
	"image/jpeg"
	"os"
)

//func CompressImage(inputPath, outputPath string, quality int) error {
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
//	// Create the output file
//	outputFile, err := os.Create(outputPath)
//	if err != nil {
//		return fmt.Errorf("failed to create output file: %v", err)
//	}
//	defer outputFile.Close()
//
//	// Encode the image with the specified quality
//	err = jpeg.Encode(outputFile, inputImage, &jpeg.Options{Quality: quality})
//	if err != nil {
//		return fmt.Errorf("failed to encode output image: %v", err)
//	}
//
//	return nil
//}

func CompressImage(inputPath, outputPath string, quality int) error {
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

	// Perform image compression
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, inputImage, &jpeg.Options{Quality: quality})
	if err != nil {
		return fmt.Errorf("failed to encode output image: %v", err)
	}

	// Update the Image entity with compressed data
	imageEntity.Data = nil // Set to nil since we don't store the image data in memory

	return nil
}
