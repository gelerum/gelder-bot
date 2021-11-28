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
func createTransactionHistory(expenses []storage.Transactions, income []storage.Transactions) string {
	history := "Expenses:\n"
	history += formatTransactions(expenses)
	history += "\nIncome:\n"
	history += formatTransactions(income)
	return history
}

func formatTransactions(transactions []storage.Transactions) string {
	var transactionList string
	for n, transaction := range transactions {
		category := transaction.Category
		amount := transaction.Amount
		creationDate := transaction.CreationDate
		transactionList += strconv.Itoa(n+1) + ". " + creationDate.Format("Jan 02") + " | " + category + " | " + strconv.FormatFloat(amount, 'f', -1, 64) + "\n"
	}
	return transactionList
}
