package telegram

import (
	"strconv"

	"github.com/gelerum/gelder-bot/pkg/storage"
)

func isCategoryValid(category string, kind string) bool {
	expensesCategories := [5]string{
		"food",
		"transportation",
		"savings",
		"others",
		"subscribtions",
	}
	incomeCategories := [5]string{
		"job",
		"freelancing",
		"buisness",
		"cashback",
		"others",
	}
	var categories [5]string
	if kind == "expenses" {
		categories = expensesCategories
	} else if kind == "income" {
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
