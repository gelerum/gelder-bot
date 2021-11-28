package storage

import (
	"strconv"
)

func createTransactionList(transactions []transactions) string {
	var history string
	for n, transaction := range transactions {
		category := transaction.Category
		amount := transaction.Amount
		creationDate := transaction.CreationDate
		history += strconv.Itoa(n+1) + ". " + creationDate.Format("Jan 02") + " | " + category + " | " + strconv.FormatFloat(amount, 'f', -1, 64) + "\n"
	}
	return history
}
