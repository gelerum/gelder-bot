package telegram

import (
	"os"

	"github.com/gelerum/gelder-bot/pkg/storage"
	tb "gopkg.in/tucnak/telebot.v2"
)

type bot struct {
	Bot    *tb.Bot
	client *storage.Client
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
	client, err := storage.NewClient()
	if err != nil {
		return nil, err
	}
	return &bot{
		Bot:    newBot,
		client: client,
	}, nil
}
