package storage

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func NewClient() (*Client, error) {
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
	coll := newClient.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("DATABASE_COLLECTION"))
	return &Client{
		Client: newClient,
		Coll:   coll,
	}, nil
}
