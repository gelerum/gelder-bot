package storage

import (
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type transactions struct {
	Category     string    `bson:"category"`
	Amount       float64   `bson:"amount"`
	CreationDate time.Time `bson:"creationDate"`
}

type user struct {
	ID       primitive.ObjectID `bson:"_id"`
	ChatID   int                `bson:"chatID"`
	Expenses []transactions     `bson:"expenses"`
	Income   []transactions     `bson:"income"`
}

func a(transactions []transactions) string {
	var history string
	for _, transaction := range transactions {
		category := transaction.Category
		amount := transaction.Amount
		creationDate := transaction.CreationDate
		history += "- " + creationDate.Format("Jan 02") + " | " + category + " | " + strconv.FormatFloat(amount, 'f', -1, 64) + "\n"
	}
	return history
}

func (c client) GetTransactions(chatID int) string {
	var doc user
	filter := bson.M{"chatID": chatID}

	c.Coll.FindOne(context.TODO(), filter).Decode(&doc)
	history := "Expenses:\n"
	history += a(doc.Expenses)
	history += "\nIncome:\n"
	history += a(doc.Income)
	return history
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
	update := bson.M{"$push": bson.M{kind: bson.D{{"category", category}, {"amount", amount}, {"creationDate", time.Now()}}}}
	_, err := c.Coll.UpdateOne(context.TODO(), filter, update)
	return err
}
