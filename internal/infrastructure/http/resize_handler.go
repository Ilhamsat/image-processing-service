package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"image-processing/internal/application"
	"net/http"
	"strconv"
)

func ResizeHandler(c *fiber.Ctx) error {
	// Parse request parameters
	inputPath := c.Query("input")
	outputPath := c.Query("output")
	widthStr := c.Query("width")
	heightStr := c.Query("height")

	// Convert width and height to integers
	width, err := strconv.Atoi(widthStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid width")
	}
	height, err := strconv.Atoi(heightStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid height")
	}

	// Call application layer to perform image resizing
	err = application.ResizeImage(inputPath, outputPath, width, height)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Failed to resize image: %v", err))
	}

	// Send success response
	return c.SendString("Image resized successfully")
}
