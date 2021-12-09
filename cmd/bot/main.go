package main

import (
	"log"
	"os"

	"github.com/gelerum/gelder-bot/internal/config"
	"github.com/gelerum/gelder-bot/internal/storage"
	"github.com/gelerum/gelder-bot/internal/telegram"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var cfg config.Config
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = config.ReadConfig(path+"/configs", &cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	config.InitBotEnvVars(&cfg)
	config.InitClientEnvVars(&cfg)
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
