package storage

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type client struct {
	client *mongo.Client
}

func newClient() *client {
	mongoClient, _ := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient.Connect(mongoContext)
	return &client{
		client: mongoClient,
	}
}

var c *client

func InitClient() {
	c = newClient()
}

func GetClient() *client {
	return c
}
