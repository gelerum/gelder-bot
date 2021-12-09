package telegram

import (
	"os"
	"testing"

	"github.com/gelerum/gelder-bot/internal/config"
	"github.com/gelerum/gelder-bot/internal/storage"
)

func TestNewBot(t *testing.T) {
	var cfg config.Config
	path, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}
	err = config.ReadConfig(path+"/configs", &cfg)
	if err != nil {
		t.Error(err)
		return
	}
	config.InitBotEnvVars(&cfg)
	config.InitClientEnvVars(&cfg)
	c, err := storage.NewClient(&cfg.Client)
	if err != nil {
		t.Error(err)
	}
	NewBot(&cfg.Bot, &cfg.Messages, c)
}
