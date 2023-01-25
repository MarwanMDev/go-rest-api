package handlers

import (
	"context"

	"github.com/MarwanMDev/go-rest-api/database"
	"github.com/MarwanMDev/go-rest-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type LibraryDTO struct {
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
}

func GetLibraries(c *fiber.Ctx) error {
	libraryCollection := database.GetCollection("libraries")
	cursor, err := libraryCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return err
	}

	var libraries []models.Library
	if err = cursor.All(context.TODO(), &libraries); err != nil {
		return err
	}

	return c.JSON(libraries)
}

func CreateLibrary(c *fiber.Ctx) error {
	newLibrary := new(LibraryDTO)

	if err := c.BodyParser(newLibrary); err != nil {
		return err
	}

	libraryCollection := database.GetCollection("libraries")
	nDoc, err := libraryCollection.InsertOne(context.TODO(), newLibrary)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"id": nDoc.InsertedID})
}
