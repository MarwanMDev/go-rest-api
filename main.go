package main

import (
	"os"

	"github.com/MarwanMDev/go-rest-api/database"
	"github.com/MarwanMDev/go-rest-api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// app := fiber.New()
	app := generateApp()
	// initialize app
	err := initializeApp()
	if err != nil {
		panic(err)
	}

	defer database.CloseMongoDB()

	// app.Post("/", func(c *fiber.Ctx) error {
	// 	sampleDoc := bson.M{"name": "Sample Todo"}
	// 	collection := database.GetCollection("todos")

	// 	nDoc, err := collection.InsertOne(context.TODO(), sampleDoc)
	// 	if err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	// 	}

	// 	return c.JSON(nDoc)
	// })

	// Get port from environment variable
	port := os.Getenv("PORT")

	app.Listen(":" + port)
}

func generateApp() *fiber.App {
	// Create fiber App
	app := fiber.New()

	// Create health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Create library group & routes
	libraryGroup := app.Group("/library")
	libraryGroup.Get("/", handlers.GetLibraries)
	libraryGroup.Post("/", handlers.CreateLibrary)
	libraryGroup.Post("/book", handlers.CreateBook)

	return app
}

func initializeApp() error {
	// setup environment variables
	err := loadEnv()
	if err != nil {
		return err
	}

	err = database.StartMongoDB()
	if err != nil {
		return err
	}

	return nil
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
