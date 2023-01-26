package handlers

import (
	"context"

	"github.com/MarwanMDev/go-rest-api/database"
	"github.com/MarwanMDev/go-rest-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type BookDTO struct {
	Title     string `json:"title" bson:"title"`
	Author    string `json:"author" bson:"author"`
	ISBN      string `json:"isbn" bson:"isbn"`
	LibraryId string `json:"library_id" bson:"library_id"`
}

func CreateBook(c *fiber.Ctx) error {
	newBook := new(BookDTO)

	if err := c.BodyParser(newBook); err != nil {
		return err
	}

	booksCollection := database.GetCollection("books")
	filter := bson.D{{Key: "_id", Value: newBook.LibraryId}}

	newBookData := models.Book{
		Title:  newBook.Title,
		Author: newBook.Author,
		ISBN:   newBook.ISBN,
	}
	updatePayLoad := bson.D{{Key: "$push", Value: bson.M{"books": newBookData}}}

	_, err := booksCollection.UpdateOne(context.TODO(), filter, updatePayLoad)
	if err != nil {
		return err
	}
	return c.SendString("Book Created Successfuly")
}
