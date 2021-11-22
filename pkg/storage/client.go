package storage

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type client struct {
	Client *mongo.Client
}

func NewClient() (*client, error) {
	newClient, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return nil, err
	}
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = newClient.Connect(mongoContext)
	if err != nil {
		return nil, err
	}
	return &client{
		Client: newClient,
	}, nil
}
