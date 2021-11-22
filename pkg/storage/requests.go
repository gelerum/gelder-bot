package storage

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func (c client) CreateUserDocument(chatID int) error {
	coll := c.Client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("DATABASE_COLLECTION"))
	count, err := coll.CountDocuments(context.TODO(), bson.D{{"chatID", chatID}})
	if err != nil {
		return err
	}
	if count != 1 {
		document := bson.D{{"chatID", chatID}, {"expenses", bson.A{}}, {"income", bson.A{}}}
		_, err = coll.InsertOne(context.TODO(), document)
	}
	return err
}

func (c client) AddTransaction(chatID int, amount float64, category string, kind string) error {
	coll := c.Client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("DATABASE_COLLECTION"))
	filter := bson.D{{"chatID", chatID}}
	update := bson.M{"$push": bson.M{kind: bson.D{{"category", category}, {"amount", amount}}}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}
