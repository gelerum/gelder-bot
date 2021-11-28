package telegram

import (
	"log"
	"strconv"
	"strings"

	"github.com/gelerum/gelder-bot/pkg/util"
	tb "gopkg.in/tucnak/telebot.v2"
)

func (b *bot) HandleStart(m *tb.Message) {
	_, err := b.Bot.Send(m.Sender, "Gelder bot helps you organize your expenses and income\n\n/help - show help")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = b.client.CreateUserDocument(m.Sender.ID)
	if err != nil {
		log.Fatal(err)
	}
}
func (b *bot) HandleHelp(m *tb.Message) {
	_, err := b.Bot.Send(m.Sender, "/help - show help \n/categories - show all categories\n\nSend me a message to add expense or income\nExample: 150.99 food expense")
	if err != nil {
		log.Fatal(err)
		return
	}
}
func (b *bot) HandleCategories(m *tb.Message) {
	_, err := b.Bot.Send(m.Sender, "Expenses:\n    1. Food\n    2. Transportation\n    3. Savings\n    4. Subscribtions\n    5. Others\n\nIncome: \n    1. Job\n    2. Freelancing\n    3. Buisness\n    4. Cashback\n    5. Others")
	if err != nil {
		log.Fatal(err)
		return
	}
}
func (b *bot) HandleGetTransactions(m *tb.Message) {
	output := b.client.GetTransactions(m.Sender.ID)
	b.Bot.Send(m.Sender, output)
}
func (b *bot) HandleMessage(m *tb.Message) {
	amountCategoryKind := strings.Fields(m.Text)
	if len(amountCategoryKind) != 3 {
		b.Bot.Send(m.Sender, "Send me amount, category and type. Example: 150.99 food expenses")
		return
	}
	amount, err := strconv.ParseFloat(amountCategoryKind[0], 64)
	if err != nil {
		b.Bot.Send(m.Sender, "Send me amount, category and type. Example: 150.99 food expenses")
		return
	}
	category := amountCategoryKind[1]
	kind := strings.ToLower(amountCategoryKind[2])
	if kind == "expenses" && util.IsCategoryValid(category, kind) {
		err = b.client.AddTransaction(m.Sender.ID, amount, category, kind)
		if err != nil {
			log.Fatal(err)
			return
		}
	} else if kind == "income" && util.IsCategoryValid(category, kind) {
		err = b.client.AddTransaction(m.Sender.ID, amount, category, kind)
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		b.Bot.Send(m.Sender, "Use expense or income. Example: 150.99 food expenses. /categories - list of supported categories")
	}
}
