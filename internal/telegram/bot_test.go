package telegram

import (
	"os"
	"testing"

	"github.com/gelerum/gelder-bot/internal/config"
	"github.com/gelerum/gelder-bot/internal/storage"
)

func TestNewBot(t *testing.T) {
	var cfg config.Config
	err := config.ReadConfig(os.Getenv("INITIAL_DIRECTORY")+"/configs", &cfg)
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
