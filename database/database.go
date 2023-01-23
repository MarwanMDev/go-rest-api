package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func StartMongoDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	var err error
	MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return nil
}

func CloseMongoDB() {
	if err := MongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func GetCollection(name string) *mongo.Collection {
	return (*mongo.Collection)(MongoClient.Database("test_go").Collection("users"))
}
