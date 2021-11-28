package util

import (
	"strconv"

	"github.com/gelerum/gelder-bot/pkg/storage"
)

func CreateTransactionList(transactions []storage.Transactions) string {
	var history string
	for n, transaction := range transactions {
		category := transaction.Category
		amount := transaction.Amount
		creationDate := transaction.CreationDate
		history += strconv.Itoa(n+1) + ". " + creationDate.Format("Jan 02") + " | " + category + " | " + strconv.FormatFloat(amount, 'f', -1, 64) + "\n"
	}
	return history
}

func IsCategoryValid(category string, kind string) bool {
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
	} else {
		categories = incomeCategories
	}
	for _, item := range categories {
		if category == item {
			return true
		}
	}
	return false
}
