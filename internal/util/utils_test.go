package util

import (
	"testing"
	"time"

	"github.com/gelerum/gelder-bot/internal/storage"
)

func TestCategoryIsValid(t *testing.T) {
	if CategoryIsValid("food", "expenses") != true {
		t.Error(`CategoryIsValid("food", "expenses") = false`)
	}
	if CategoryIsValid("job", "income") != true {
		t.Error(`CategoryIsValid("job", "income") = false`)
	}
	if CategoryIsValid("food", "income") != false {
		t.Error(`CategoryIsValid("food", "income") = true`)
	}
	if CategoryIsValid("job", "expenses") != false {
		t.Error(`CategoryIsValid("job", "expenses") = true`)
	}
}

func TestCalculateTransactionsSum(t *testing.T) {
	firstDate, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2021-12-08 13:56:55.362 +0000 UTC")
	firstTransaction := storage.Transaction{
		Category:     "food",
		Amount:       10,
		CreationDate: firstDate,
	}
	secondDate, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2021-12-08 14:00:56.314 +0000 UTC")
	secondTransaction := storage.Transaction{
		Category:     "food",
		Amount:       20,
		CreationDate: secondDate,
	}
	transactions := []storage.Transaction{firstTransaction, secondTransaction}
	sum := CalculateTransactionsSum(transactions)
	if sum != 30 {
		t.Error("CalculateTransactionsSum([]storage.Transaction{firstTransaction, secondTransaction}) =", sum)
	}
}

func TestCreateTransactionHistory(t *testing.T) {
	firstDate, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2021-12-08 13:56:55.362 +0000 UTC")
	firstExpense := storage.Transaction{
		Category:     "food",
		Amount:       10,
		CreationDate: firstDate,
	}
	secondDate, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2021-12-08 14:00:56.314 +0000 UTC")
	secondExpense := storage.Transaction{
		Category:     "others",
		Amount:       2,
		CreationDate: secondDate,
	}
	transactionHistory := CreateTransactionHistory([]storage.Transaction{firstExpense, secondExpense})
	if transactionHistory != "1. Dec 08 | food | 10\n2. Dec 08 | others | 2\n" {
		t.Error("transactionHistory([]storage.Transaction{firstTransaction, secondTransaction}) =", transactionHistory)
	}
}
