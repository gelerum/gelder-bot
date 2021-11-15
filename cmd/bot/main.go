package main

import (
	"github.com/gelerum/gelder-bot/pkg/storage"
	"github.com/gelerum/gelder-bot/pkg/telegram"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	telegram.InitBot()
	storage.InitClient()
	telegram.GetBot().Bot.Handle("/start", telegram.GetBot().HandleStart)
	telegram.GetBot().Bot.Handle("/help", telegram.GetBot().HandleHelp)
	telegram.GetBot().Bot.Handle("/categories", telegram.GetBot().HandleCategories)
	telegram.GetBot().Bot.Handle(tb.OnText, telegram.GetBot().HandleMessage)
	telegram.GetBot().Bot.Start()
}