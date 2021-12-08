package main

import (
	"log"

	"github.com/gelerum/gelder-bot/internal/config"
	"github.com/gelerum/gelder-bot/internal/storage"
	"github.com/gelerum/gelder-bot/internal/telegram"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	client, err := storage.NewClient(&cfg.Client)
	if err != nil {
		log.Fatal(err)
		return
	}
	bot, err := telegram.NewBot(&cfg.Bot, &cfg.Messages, client)
	if err != nil {
		log.Fatal(err)
		return
	}
	bot.Bot.Handle("/start", bot.HandleStart)
	bot.Bot.Handle("/help", bot.HandleHelp)
	bot.Bot.Handle("/categories", bot.HandleCategories)
	bot.Bot.Handle("/balance", bot.HandleBalance)
	bot.Bot.Handle("/transactions", bot.HandleTransactions)
	bot.Bot.Handle("/del", bot.DeleteTransactions)
	bot.Bot.Handle(tb.OnText, bot.HandleMessage)
	bot.Bot.Start()
}
