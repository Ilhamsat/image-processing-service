package main

import (
	"github.com/gofiber/fiber/v2"
	"image-processing/internal/infrastructure/http"
)

func main() {
	app := fiber.New()

	// Register routes
	app.Post("/convert", http.ConvertHandler)
	app.Post("/resize", http.ResizeHandler)
	app.Post("/compress", http.CompressHandler)

	// Start the server
	err := app.Listen(":3003")
	if err != nil {
		panic(err)
	}
}
