package domain

import "errors"

var (
	ErrImageNotFound    = errors.New("image not found")
	ErrInvalidImageType = errors.New("invalid image type")
	ErrImageProcessing  = errors.New("image processing failed")
)
