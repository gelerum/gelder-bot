package storage

import (
	"os"
	"testing"
	"time"

	"github.com/gelerum/gelder-bot/internal/config"
)

func TestGetTransactions(t *testing.T) {
	cfg := config.Client{
		URI:        os.Getenv("MONGO_URI"),
		Name:       os.Getenv("DATABASE_NAME"),
		Collection: os.Getenv("DATABASE_TEST_COLLECTION"),
	}
	client, _ := NewClient(&cfg)
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
	expenses, income := client.GetTransactions(1)
	if expenses[0] != expectedExpense || income[0] != expectedIncome {
		t.Error("GetTransactions(1) =", expenses, income)
	}
	expenses, income = client.GetTransactions(0)
	if expenses != nil || income != nil {
		t.Error("GetTransactions(0) =", expenses, income)
	}
}