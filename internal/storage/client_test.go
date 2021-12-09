package storage

import (
	"os"
	"testing"

	"github.com/gelerum/gelder-bot/internal/config"
)

func TestNewClient(t *testing.T) {
	cfg := config.Client{
		URI:        os.Getenv("MONGO_URI"),
		Name:       os.Getenv("DATABASE_NAME"),
		Collection: os.Getenv("DATABASE_TEST_COLLECTION"),
	}
	_, err := NewClient(&cfg)
	if err != nil {
		t.Error(cfg.URI, cfg.Name, cfg.Collection)
	}
}
