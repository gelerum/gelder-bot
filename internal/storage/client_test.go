package storage

import (
	"testing"

	"github.com/gelerum/gelder-bot/internal/config"
)

func TestNewClient(t *testing.T) {
	var cfg config.Config
	config.InitClientEnvVars(&cfg)
	_, err := NewClient(&cfg.Client)
	if err != nil {
		t.Error(err)
	}
}
