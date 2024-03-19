package domain

import "image"

type Image struct {
	ID       string
	Format   string
	Data     []byte
	Width    int
	Height   int
	Metadata map[string]string
}

func NewImage(id, format string, data []byte, img image.Image, metadata map[string]string) *Image {
	return &Image{
		ID:       id,
		Format:   format,
		Data:     data,
		Width:    img.Bounds().Dx(),
		Height:   img.Bounds().Dy(),
		Metadata: metadata,
	}
}
