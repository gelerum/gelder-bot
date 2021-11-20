package telegram

import (
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

type bot struct {
	Bot *tb.Bot
}

func NewBot() (*bot, error) {
	var (
		port  = os.Getenv("PORT")
		url   = os.Getenv("URL")
		token = os.Getenv("TOKEN")
	)
	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: url},
	}
	preferencies := tb.Settings{
		Token:  token,
		Poller: webhook,
	}
	newBot, err := tb.NewBot(preferencies)
	if err != nil {
		return nil, err
	}
	return &bot{
		Bot: newBot,
	}, nil
}
