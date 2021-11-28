package telegram

import (
	"log"
	"strconv"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

func (b *bot) HandleStart(m *tb.Message) {
	_, err := b.Bot.Send(m.Sender, b.messages.Start)
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
	_, err := b.Bot.Send(m.Sender, b.messages.Help)
	if err != nil {
		log.Fatal(err)
		return
	}
}
func (b *bot) HandleCategories(m *tb.Message) {
	_, err := b.Bot.Send(m.Sender, b.messages.Categories)
	if err != nil {
		log.Fatal(err)
		return
	}
}
func (b *bot) HandleTransactions(m *tb.Message) {
	output := b.client.GetTransactions(m.Sender.ID)
	b.Bot.Send(m.Sender, output)
}
func (b *bot) HandleMessage(m *tb.Message) {
	amountCategoryKind := strings.Fields(m.Text)
	if len(amountCategoryKind) != 3 {
		b.Bot.Send(m.Sender, b.messages.InitialError)
		return
	}
	amount, err := strconv.ParseFloat(amountCategoryKind[0], 64)
	if err != nil {
		b.Bot.Send(m.Sender, b.messages.InitialError)
		return
	}
	category := amountCategoryKind[1]
	kind := strings.ToLower(amountCategoryKind[2])
	if kind == "expenses" && isCategoryValid(category, kind) {
		err = b.client.AddTransaction(m.Sender.ID, amount, category, kind)
		if err != nil {
			log.Fatal(err)
			return
		}
	} else if kind == "income" && isCategoryValid(category, kind) {
		err = b.client.AddTransaction(m.Sender.ID, amount, category, kind)
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		b.Bot.Send(m.Sender, b.messages.CategoryError)
	}
}
