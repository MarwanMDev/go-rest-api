package main

import (
	"context"

	"github.com/MarwanMDev/go-rest-api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	app := fiber.New()
	// initialize app
	err := initializeApp()
	if err != nil {
		panic(err)
	}

	defer database.CloseMongoDB()

	app.Post("/", func(c *fiber.Ctx) error {
		sampleDoc := bson.M{"name": "Sample Todo"}
		collection := database.GetCollection("todos")

		nDoc, err := collection.InsertOne(context.TODO(), sampleDoc)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(nDoc)
	})

	app.Listen(":3000")
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
