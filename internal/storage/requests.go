package storage

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// Expense or income transaction structure
	Transaction struct {
		Category     string    `bson:"category"`
		Amount       float64   `bson:"amount"`
		CreationDate time.Time `bson:"creationDate"`
	}
	// User structure
	user struct {
		ID       primitive.ObjectID `bson:"_id"`
		ChatID   int                `bson:"chatID"`
		Expenses []Transaction      `bson:"expenses"`
		Income   []Transaction      `bson:"income"`
	}
)

// Get user's expenses and income transactions
func (c *Client) GetTransactions(chatID int) ([]Transaction, []Transaction) {
	var doc user
	filter := bson.M{"chatID": chatID}
	c.coll.FindOne(c.ctx, filter).Decode(&doc)
	return doc.Expenses, doc.Income
}

// Create user document if it doesn't exist
func (c *Client) CreateUserDocument(chatID int) error {
	count, err := c.coll.CountDocuments(c.ctx, bson.M{"chatID": chatID})
	if err != nil {
		return err
	}
	if count == 1 {
		return nil
	}
	document := bson.D{{"chatID", chatID}, {"expenses", bson.A{}}, {"income", bson.A{}}}
	_, err = c.coll.InsertOne(c.ctx, document)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Print("User has been added successfully")
	return nil
}

// Add expense or income transaction
func (c *Client) AddTransaction(chatID int, amount float64, category string, kind string) error {
	filter := bson.M{"chatID": chatID}
	update := bson.M{"$push": bson.M{kind: bson.D{{"category", category}, {"amount", amount}, {"creationDate", time.Now()}}}}
	_, err := c.coll.UpdateOne(c.ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// Delete expense or income transaction
func (c *Client) DeleteTransaction(chatID int, transactionNumber string, kind string) error {
	filter := bson.M{"chatID": chatID}
	unsetUpdate := bson.M{"$unset": bson.D{{kind + "." + transactionNumber, 1}}}
	pullUpdate := bson.M{"$pull": bson.M{kind: nil}}
	_, err := c.coll.UpdateOne(c.ctx, filter, unsetUpdate)
	if err != nil {
		return err
	}
	_, err = c.coll.UpdateOne(c.ctx, filter, pullUpdate)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
