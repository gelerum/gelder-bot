package telegram

import (
	"log"

	"github.com/gelerum/gelder-bot/internal/config"
	"github.com/gelerum/gelder-bot/internal/storage"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Telegram bot structure
type bot struct {
	Bot      *tb.Bot
	client   *storage.Client
	config   *config.Bot
	messages *config.Messages
}

// Create new bot. Set configs and client
func NewBot(cfg *config.Bot, msgs *config.Messages, clnt *storage.Client) (*bot, error) {
	var (
		port     = cfg.Port
		appUrl   = cfg.AppURL
		botToken = cfg.Token
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
		log.Fatal(err)
		return nil, err
	}

	return &bot{
		Bot:      newBot,
		client:   clnt,
		config:   cfg,
		messages: msgs,
	}, nil
}
