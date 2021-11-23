package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	ChatID   int                `bson:"chatID" json:"chatID"`
	Expenses []interface{}      `bson:"expenses" json:"expenses"`
	Income   []interface{}      `bson:"income" json:"income,"`
}

func (c client) Get(chatID int) bson.M {
	fmt.Print(chatID)
	filter := bson.M{"chatID": chatID}
	var a bson.M
	c.Coll.FindOne(context.TODO(), filter).Decode(&a)
	fmt.Printf("Found a single document: %+v\n", a)
	return a
}

func (c client) CreateUserDocument(chatID int) error {
	count, err := c.Coll.CountDocuments(context.TODO(), bson.D{{"chatID", chatID}})
	if err != nil {
		return err
	}
	if count != 1 {
		document := bson.D{{"chatID", chatID}, {"expenses", bson.A{}}, {"income", bson.A{}}}
		_, err = c.Coll.InsertOne(context.TODO(), document)
	}
	return err
}

func (c client) AddTransaction(chatID int, amount float64, category string, kind string) error {
	filter := bson.D{{"chatID", chatID}}
	update := bson.M{"$push": bson.M{kind: bson.D{{"category", category}, {"amount", amount}}}}
	_, err := c.Coll.UpdateOne(context.TODO(), filter, update)
	return err
}
