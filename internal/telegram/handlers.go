package telegram

import (
	"log"
	"strconv"
	"strings"

	"github.com/gelerum/gelder-bot/internal/util"
	tb "gopkg.in/tucnak/telebot.v2"
)

func HandleStart(b *bot, m *tb.Message) {
	_, err := b.Bot.Send(m.Sender, b.messages.Start)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = b.client.CreateUserDocument(m.Sender.ID)
	if err != nil {
		log.Fatal(err)
		b.Bot.Send(m.Sender, "Bot error. Try /start again")
		return
	}
}

func (b *bot) HandleHelp(m *tb.Message) {
	_, err := b.Bot.Send(m.Sender, b.messages.Help)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Print("Help has been sent successfully")
}

func (b *bot) HandleCategories(m *tb.Message) {
	_, err := b.Bot.Send(m.Sender, b.messages.Categories)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Print("Categories list has been sent successfully")
}

func (b *bot) HandleBalance(m *tb.Message) {
	expenses, income := b.client.GetTransactions(m.Sender.ID)
	expensesSum := util.CalculateTransactionsSum(expenses)
	incomeSum := util.CalculateTransactionsSum(income)
	output := "Expenses: " + strconv.FormatFloat(expensesSum, 'f', -1, 64) + "\nIncome: " + strconv.FormatFloat(incomeSum, 'f', -1, 64)
	_, err := b.Bot.Send(m.Sender, output)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Print("Balance has been sent successfully")
}

func (b *bot) HandleTransactions(m *tb.Message) {
	expenses, income := b.client.GetTransactions(m.Sender.ID)
	output := "Expenses:\n"
	output += util.CreateTransactionHistory(expenses)
	output += "\nIncome:\n"
	output += util.CreateTransactionHistory(income)
	log.Print("Transactions list has been sent successfully")
	b.Bot.Send(m.Sender, output)
}

func (b *bot) DeleteTransactions(m *tb.Message) {
	numberKind := strings.Fields(m.Text)
	if len(numberKind) != 3 {
		b.Bot.Send(m.Sender, b.messages.DelInitialError)
		return
	}
	kind := strings.ToLower(numberKind[2])
	initialTransactionNumber, err := strconv.Atoi(numberKind[1])
	if err != nil {
		b.Bot.Send(m.Sender, b.messages.NumberError)
		return
	}
	if initialTransactionNumber < 1 {
		b.Bot.Send(m.Sender, b.messages.NumberError)
		return
	}
	transactionNumber := strconv.Itoa(initialTransactionNumber - 1)
	err = b.client.DeleteTransaction(m.Sender.ID, transactionNumber, kind)
	if err != nil {
		log.Fatal(err)
		b.Bot.Send(m.Sender, "Bot error")
		return
	}
	log.Print("Transaction has been deleted successfully")
	b.Bot.Send(m.Sender, "Transaction has been deleted successfully")
}

func (b *bot) HandleMessage(m *tb.Message) {
	amountCategoryKind := strings.Fields(m.Text)
	if len(amountCategoryKind) != 3 {
		b.Bot.Send(m.Sender, b.messages.AddInitialError)
		return
	}
	amount, err := strconv.ParseFloat(amountCategoryKind[0], 64)
	if err != nil {
		b.Bot.Send(m.Sender, b.messages.AddInitialError)
		return
	}
	category := strings.ToLower(amountCategoryKind[1])
	kind := strings.ToLower(amountCategoryKind[2])
	if kind != "expenses" && kind != "income" {
		b.Bot.Send(m.Sender, b.messages.KindError)
		return
	}
	if !util.CategoryIsValid(category, kind) {
		b.Bot.Send(m.Sender, b.messages.CategoryError)
		return
	}
	err = b.client.AddTransaction(m.Sender.ID, amount, category, kind)
	if err != nil {
		log.Fatal(err)
		b.Bot.Send(m.Sender, "Bot error")
		return
	}
	log.Print("Transaction has been added successfully")
	b.Bot.Send(m.Sender, "Transaction has been added successfully")
}
