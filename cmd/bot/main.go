package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/gelerum/gelder-bot/pkg/storage"
	"github.com/gelerum/gelder-bot/pkg/telegram"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	client, err := storage.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}
	bot, err := telegram.NewBot()
	if err != nil {
		log.Fatal(err)
		return
	}
	bot.Bot.Handle("/start", func(m *tb.Message) {
		_, err := bot.Bot.Send(m.Sender, "Gelder bot helps you organize your expenses and income\n\n/help - show help")
		if err != nil {
			log.Fatal(err)
			return
		}
		err = client.CreateUserDocument(m.Sender.ID)
		if err != nil {
			log.Fatal(err)
		}
	})
	bot.Bot.Handle("/help", func(m *tb.Message) {
		_, err := bot.Bot.Send(m.Sender, "/help - show help \n/categories - show all categories\n\nSend me a message to add expense or income\nExample: 150.99 food expense")
		if err != nil {
			log.Fatal(err)
			return
		}
	})
	bot.Bot.Handle("/categories", func(m *tb.Message) {
		_, err := bot.Bot.Send(m.Sender, "Expenses:\n    1. Food\n    2. Transportation\n    3. Savings\n    4. Subscribtions\n    5. Others\n\nIncome: \n    1. Job\n    2. Freelancing\n    3. Buisness\n    4. Cashback\n    5. Others")
		if err != nil {
			log.Fatal(err)
			return
		}
	})
	bot.Bot.Handle("/get", func(m *tb.Message) {
		// pass
	})
	bot.Bot.Handle(tb.OnText, func(m *tb.Message) {
		amountCategoryKind := strings.Fields(m.Text)
		if len(amountCategoryKind) != 3 {
			bot.Bot.Send(m.Sender, "Send me amount, category and type. Example: 150.99 food expenses")
			return
		}
		amount, err := strconv.ParseFloat(amountCategoryKind[0], 64)
		if err != nil {
			bot.Bot.Send(m.Sender, "Send me amount, category and type. Example: 150.99 food expenses")
			return
		}
		category := amountCategoryKind[1]
		kind := strings.ToLower(amountCategoryKind[2])
		if kind == "expenses" && isCategoryValid(category, kind) {
			err = client.AddTransaction(m.Sender.ID, amount, category, kind)
			if err != nil {
				log.Fatal(err)
				return
			}
		} else if kind == "income" && isCategoryValid(category, kind) {
			err = client.AddTransaction(m.Sender.ID, amount, category, kind)
			if err != nil {
				log.Fatal(err)
				return
			}
		} else {
			bot.Bot.Send(m.Sender, "Use expense or income. Example: 150.99 food expenses. /categories - list of supported categories")
		}
	})
}
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
