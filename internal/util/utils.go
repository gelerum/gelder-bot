package util

import (
	"strconv"

	"github.com/gelerum/gelder-bot/internal/storage"
)

func CategoryIsValid(category string, kind string) bool {
	expensesCategories := []string{
		"food",
		"transportation",
		"entertainment",
		"household",
		"savings",
		"others",
		"subscribtions",
	}
	incomeCategories := []string{
		"job",
		"freelancing",
		"buisness",
		"cashback",
		"others",
	}
	var categories []string
	if kind == "expenses" {
		categories = expensesCategories
	}
	if kind == "income" {
		categories = incomeCategories
	}
	for _, item := range categories {
		if category == item {
			return true
		}
	}
	return false
}

func CalculateTransactionsSum(transactions []storage.Transaction) float64 {
	var sum float64
	for _, transaction := range transactions {
		sum += transaction.Amount
	}
	return sum
}

func CreateTransactionHistory(transactions []storage.Transaction) string {
	var transactionHistory string
	for n, transaction := range transactions {
		category := transaction.Category
		amount := transaction.Amount
		creationDate := transaction.CreationDate
		transactionHistory += strconv.Itoa(n+1) + ". " + creationDate.Format("Jan 02") + " | " + category + " | " + strconv.FormatFloat(amount, 'f', -1, 64) + "\n"
	}
	return transactionHistory
}
