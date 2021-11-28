package storage

import (
	"context"
	"time"

	"github.com/gelerum/gelder-bot/pkg/util"
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

func (c Client) GetTransactions(chatID int) string {
	var doc user
	filter := bson.M{"chatID": chatID}
	c.Coll.FindOne(context.TODO(), filter).Decode(&doc)

	history := "Expenses:\n"
	history += util.CreateTransactionList(doc.Expenses)
	history += "\nIncome:\n"
	history += util.CreateTransactionList(doc.Income)
	return history
}

func (c Client) CreateUserDocument(chatID int) error {
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

func (c Client) AddTransaction(chatID int, amount float64, category string, kind string) error {
	filter := bson.D{{"chatID", chatID}}
	update := bson.M{"$push": bson.M{kind: bson.D{{"category", category}, {"amount", amount}, {"creationDate", time.Now()}}}}
	_, err := c.Coll.UpdateOne(context.TODO(), filter, update)
	return err
}
