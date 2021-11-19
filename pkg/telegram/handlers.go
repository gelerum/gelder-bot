package telegram

import (
	"strconv"
	"strings"

	"github.com/gelerum/gelder-bot/pkg/storage"
	tb "gopkg.in/tucnak/telebot.v2"
)

func (b *bot) HandleHelp(message *tb.Message) {
	b.Bot.Send(message.Sender, "/help - show help \n/categories - show all categories\n\nSend me a message to add expense or income\nExample: 150.99 food expense")

}

func (b *bot) HandleStart(message *tb.Message) {
	b.Bot.Send(message.Sender, "Gelder bot helps you organize your expenses and income\n\n/help - show help")
	storage.GetClient().CreateUserDocument(message.Sender.ID)
}

func (b *bot) HandleCategories(message *tb.Message) {
	b.Bot.Send(message.Sender, "Expenses:\n    1. Food\n    2. Transportation\n    3. Savings\n    4. Subscribtions\n    5. Others\n\nIncome: \n    1. Job\n    2. Freelancing\n    3. Buisness\n    4. Cashback\n    5. Others")
}

func (b *bot) HandleMessage(message *tb.Message) {
	amountCategoryKind := strings.Fields(message.Text)
	if len(amountCategoryKind) != 3 {
		b.Bot.Send(message.Sender, "Send me amount, category and type. Example: 150.99 food expenses")
		return
	}
	amount, err := strconv.ParseFloat(amountCategoryKind[0], 64)
	if err != nil {
		b.Bot.Send(message.Sender, "Send me amount, category and type. Example: 150.99 food expenses")
		return
	}
	category := amountCategoryKind[1]
	kind := strings.ToLower(amountCategoryKind[2])
	if kind == "expenses" && isCategoryValid(category, kind) {
		storage.GetClient().AddPosition(message.Sender.ID, amount, category, kind)
	} else if kind == "income" && isCategoryValid(category, kind) {
		storage.GetClient().AddPosition(message.Sender.ID, amount, category, kind)
	} else {
		b.Bot.Send(message.Sender, "Use expense or income. Example: 150.99 food expenses. /categories - list of supported categories")
	}
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
