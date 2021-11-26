package storage

import (
	"context"
	"fmt"
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
	var document user
	c.Coll.FindOne(context.TODO(), filter).Decode(&document)
	sum := 0.0
	var transactions string
	for _, expense := range document.Expenses {
		category := expense.Category
		amount := expense.Amount
		creationDate := expense.CreationDate
		sum += amount
		transactions += category + strconv.FormatFloat(amount, 'f', -1, 64) + creationDate.String()
	}
	output := strconv.FormatFloat(sum, 'f', -1, 64) + "\n\n" + transactions
	return output
}

func (c client) Delete(chatID int) {
	change := bson.M{"$pull": bson.M{"expenses": bson.D{{"category", "food"}, {"amount", 11}}}}
	a, err := c.Coll.UpdateOne(context.TODO(), bson.M{"chatID": chatID}, change)
	fmt.Println(err)
	fmt.Println(a)
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
	update := bson.M{"$push": bson.M{kind: bson.D{{"category", category}, {"amount", amount}, {"created_at", time.Now()}}}}
	_, err := c.Coll.UpdateOne(context.TODO(), filter, update)
	return err
}
