package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"image-processing/internal/application"
	"net/http"
	"strconv"
)

func CompressHandler(c *fiber.Ctx) error {
	// Parse request parameters
	inputPath := c.Query("input")
	outputPath := c.Query("output")
	qualityStr := c.Query("quality")

	// Convert quality to integer
	quality, err := strconv.Atoi(qualityStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid quality")
	}

	// Call application layer to perform image compression
	err = application.CompressImage(inputPath, outputPath, quality)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Failed to compress image: %v", err))
	}

	// Send success response
	return c.SendString("Image compressed successfully")
}
