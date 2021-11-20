package storage

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func (c client) CreateUserDocument(chatID int) {
	coll := c.client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("USED_COLLECTION"))
	count, _ := coll.CountDocuments(context.TODO(), bson.D{{"chatID", chatID}})
	if count != 1 {
		document := bson.D{{"chatID", chatID}, {"expenses", bson.A{}}, {"income", bson.A{}}}
		coll.InsertOne(context.TODO(), document)
	}
}

func (c client) AddTransaction(chatID int, amount float64, category string, kind string) {
	coll := c.client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("USED_COLLECTION"))
	filter := bson.D{{"chatID", chatID}}
	update := bson.M{"$push": bson.M{kind: bson.A{bson.M{"category": category}, bson.M{"smount": amount}}}}
	coll.UpdateOne(context.TODO(), filter, update)
}

func (c client) GetTransaction(chatID int) {
	coll := c.client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("USED_COLLECTION"))
	var podcast bson.D
	coll.FindOne(context.TODO(), bson.D{{"chatID", chatID}}).Decode(&podcast)

}
