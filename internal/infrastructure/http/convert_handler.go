package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"image-processing/internal/application"
	"net/http"
)

func ConvertHandler(c *fiber.Ctx) error {
	// Parse request parameters
	inputPath := c.Query("input")
	outputPath := c.Query("output")

	// Call application layer to perform PNG to JPEG conversion
	err := application.ConvertImage(inputPath, outputPath)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Failed to convert image: %v", err))
	}

	// Send success response
	return c.SendString("Image converted successfully")
}
