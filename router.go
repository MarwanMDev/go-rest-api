package main

import (
	"github.com/MarwanMDev/go-rest-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func generateApp() *fiber.App {
	// Create fiber App
	app := fiber.New()

	// Create health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Create library group & routes
	libraryGroup := app.Group("/library")
	libraryGroup.Get("/", handlers.TestHandler)

	return app
}
