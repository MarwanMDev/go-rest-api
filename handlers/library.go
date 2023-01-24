package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type LibraryDTO struct {
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
}

func CreateLibrary(c *fiber.Ctx) error {
	newLibrary := new(LibraryDTO)

	if err := c.BodyParser(newLibrary); err != nil {
		return err
	}

	fmt.Println(newLibrary)

	return c.SendString("Library created")
}
