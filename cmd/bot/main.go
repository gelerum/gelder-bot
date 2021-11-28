package main

import (
	"log"

	"github.com/gelerum/gelder-bot/pkg/telegram"
)

func main() {

	bot, err := telegram.NewBot()
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Bot.Start()
}
