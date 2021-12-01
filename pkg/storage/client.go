package storage

import (
	"context"
	"log"
	"time"

	"github.com/gelerum/gelder-bot/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	client *mongo.Client
	coll   *mongo.Collection
	ctx    context.Context
}

func NewClient(cfg *config.Client) (*Client, error) {
	newClient, err := mongo.NewClient(options.Client().ApplyURI(cfg.URI))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = newClient.Connect(mongoContext)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	coll := newClient.Database(cfg.Name).Collection(cfg.Collection)
	context := context.TODO()
	return &Client{
		client: newClient,
		coll:   coll,
		ctx:    context,
	}, nil
}
