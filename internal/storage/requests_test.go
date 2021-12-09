package storage

import (
	"os"
	"testing"
	"time"

	"github.com/gelerum/gelder-bot/internal/config"
	"go.mongodb.org/mongo-driver/bson"
)

func initClient() (*Client, error) {
	cfg := config.Client{
		URI:        os.Getenv("MONGO_URI"),
		Name:       os.Getenv("DATABASE_NAME"),
		Collection: os.Getenv("DATABASE_TEST_COLLECTION"),
	}
	c, err := NewClient(&cfg)
	return c, err
}

func TestGetTransactions(t *testing.T) {
	c, err := initClient()
	if err != nil {
		t.Error(err)
		return
	}
	expenseDate, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2021-12-08 13:56:55.362 +0000 UTC")
	expectedExpense := Transaction{
		Category:     "food",
		Amount:       10,
		CreationDate: expenseDate,
	}
	incomeDate, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2021-12-08 14:00:56.314 +0000 UTC")
	expectedIncome := Transaction{
		Category:     "job",
		Amount:       20,
		CreationDate: incomeDate,
	}
	expenses, income := c.GetTransactions(1)
	if expenses[0] != expectedExpense || income[0] != expectedIncome {
		t.Error("GetTransactions(1) =", expenses, income)
	}
	expenses, income = c.GetTransactions(0)
	if expenses != nil || income != nil {
		t.Error("GetTransactions(0) =", expenses, income)
	}
}

func TestCreateUserDocument(t *testing.T) {
	c, err := initClient()
	if err != nil {
		t.Error(err)
		return
	}
	// first case
	err = c.CreateUserDocument(1)
	if err != nil {
		t.Error("client.CreateUserDocument(1) =", err)
		return
	}
	count, err := c.coll.CountDocuments(c.ctx, bson.M{"chatID": 1})
	if err != nil {
		t.Error(err)
		return
	}
	if count > 1 {
		t.Error("Document was added")
		return
	}
	// second case
	err = c.CreateUserDocument(2)
	if err != nil {
		t.Error("client.CreateUserDocument(2) =", err)
		return
	}
	count, err = c.coll.CountDocuments(c.ctx, bson.M{"chatID": 2})
	if err != nil {
		t.Error(err)
		return
	}
	if count == 0 {
		t.Error("Document wasn't added")
		return
	}
	// return collection to original state with delete added document
	_, err = c.coll.DeleteOne(c.ctx, bson.M{"chatID": 2})
	if err != nil {
		t.Error("Collection wasn't returned to original state with delete added document", err)
	}
}

func TestAddTransaction(t *testing.T) {
	c, err := initClient()
	if err != nil {
		t.Error(err)
		return
	}
	err = c.AddTransaction(1, 20, "food", "expenses")
	if err != nil {
		t.Error(err)
		return
	}
	err = c.AddTransaction(1, 30, "job", "income")
	if err != nil {
		t.Error(err)
		return
	}
	expectedExpense := Transaction{
		Category: "food",
		Amount:   20,
	}
	expectedIncome := Transaction{
		Category: "food",
		Amount:   30,
	}
	expenses, income := c.GetTransactions(1)
	if expenses[1] != expectedExpense {
		t.Error("expenses[1] of c.GetTransactions(1) =", expenses[1])
	}
	if income[1] != expectedIncome {
		t.Error("expenses[1] of c.GetTransactions(1) =", income[1])
	}
	// return collection to original state with delete added transactions
	err = c.DeleteTransaction(1, "1", "expenses")
	if err != nil {
		t.Error("Collection wasn't returned to original state with delete added expense", err)
	}
	err = c.DeleteTransaction(1, "1", "income")
	if err != nil {
		t.Error("Collection wasn't returned to original state with delete added income", err)
	}
}
