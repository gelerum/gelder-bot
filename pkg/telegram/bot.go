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
		port     = os.Getenv("PORT")
		appUrl   = os.Getenv("APP_URL")
		botToken = os.Getenv("BOT_TOKEN")
	)
	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: appUrl},
	}
	preferencies := tb.Settings{
		Token:  botToken,
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
