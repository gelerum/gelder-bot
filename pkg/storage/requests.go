package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transactions struct {
	Category     string    `bson:"category"`
	Amount       float64   `bson:"amount"`
	CreationDate time.Time `bson:"creationDate"`
}

type user struct {
	ID       primitive.ObjectID `bson:"_id"`
	ChatID   int                `bson:"chatID"`
	Expenses []Transactions     `bson:"expenses"`
	Income   []Transactions     `bson:"income"`
}

func (c Client) GetTransactions(chatID int) ([]Transactions, []Transactions) {
	var doc user
	filter := bson.M{"chatID": chatID}
	c.coll.FindOne(context.TODO(), filter).Decode(&doc)
	return doc.Expenses, doc.Income

}

func (c Client) CreateUserDocument(chatID int) error {
	count, err := c.coll.CountDocuments(context.TODO(), bson.D{{"chatID", chatID}})
	if err != nil {
		return err
	}
	if count != 1 {
		document := bson.D{{"chatID", chatID}, {"expenses", bson.A{}}, {"income", bson.A{}}}
		_, err = c.coll.InsertOne(context.TODO(), document)
	}
	return err
}

func (c Client) AddTransaction(chatID int, amount float64, category string, kind string) error {
	filter := bson.D{{"chatID", chatID}}
	update := bson.M{"$push": bson.M{kind: bson.D{{"category", category}, {"amount", amount}, {"creationDate", time.Now()}}}}
	_, err := c.coll.UpdateOne(context.TODO(), filter, update)
	return err
}
