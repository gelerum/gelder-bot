package storage

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func (c client) CreateUserDocument(chatID int) {
	coll := c.client.Database(os.Getenv("DATABASE_NAME")).Collection("USED_COLLECTION")
	fmt.Println(coll)
	count, _ := coll.CountDocuments(context.TODO(), bson.D{{"chatID", chatID}})
	fmt.Println(count, "my count")
	if count != 1 {
		document := bson.D{{"chatID", chatID}, {"expenses", bson.A{}}, {"income", bson.A{}}}
		_, err := coll.InsertOne(context.TODO(), document)
		fmt.Print(err)
	}
}

func (c client) AddPosition(chatID int, amount float64, category string, kind string) {
	coll := c.client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("USED_COLLECTION"))
	filter := bson.D{{"chatID", chatID}}
	update := bson.M{"$push": bson.M{kind: bson.E{category, amount}}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	fmt.Print(err, "error here")
}
