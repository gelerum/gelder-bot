package telegram

import (
	"strconv"

	"github.com/gelerum/gelder-bot/pkg/storage"
)

func categoryIsValid(category string, kind string) bool {
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
	if kind == "expenses" {
		categories = incomeCategories
	}
	for _, item := range categories {
		if category == item {
			return true
		}
	}
	return false
}

func calculateTransactionsSum(transactions []storage.Transactions) float64 {
	var sum float64
	for _, transaction := range transactions {
		sum += transaction.Amount
	}
	return sum
}

func createTransactionHistory(transactions []storage.Transactions) string {
	var transactionList string
	for n, transaction := range transactions {
		category := transaction.Category
		amount := transaction.Amount
		creationDate := transaction.CreationDate
		transactionList += strconv.Itoa(n+1) + ". " + creationDate.Format("Jan 02") + " | " + category + " | " + strconv.FormatFloat(amount, 'f', -1, 64) + "\n"
	}
	return transactionList
}
