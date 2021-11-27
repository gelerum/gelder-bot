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

func (c client) GetTransactions(chatID int) string {
	filter := bson.M{"chatID": chatID}
	var doc user
	c.Coll.FindOne(context.TODO(), filter).Decode(&doc)
	var sum float64
	var transactions string
	for _, expense := range doc.Expenses {
		category := expense.Category
		amount := expense.Amount
		creationDate := expense.CreationDate
		sum += amount
		transactions += strconv.FormatFloat(amount, 'f', -1, 64) + " " + category + " " + creationDate.Format("Jan 02") + "\n"
	}
	output := strconv.FormatFloat(sum, 'f', -1, 64) + "\n\n" + transactions
	return output
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
