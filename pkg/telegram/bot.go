package telegram

import (
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

type bot struct {
	Bot *tb.Bot
}

func newBot() *bot {
	var (
		port  = os.Getenv("WEBHOOKS_PORT")
		url   = os.Getenv("WEBHOOKS_URL")
		token = os.Getenv("BOT_TOKEN")
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: url},
	}

	preferencies := tb.Settings{
		Token:  token,
		Poller: webhook,
	}
	newBot, _ := tb.NewBot(preferencies)
	return &bot{
		Bot: newBot,
	}
}

var b *bot

func GetBot() *bot {
	if b != nil {
		return b
	}

	b = newBot()

	return b
}
