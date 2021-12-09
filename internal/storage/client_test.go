package storage

import (
	"fmt"
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
	fmt.Println("fdafjaoidoihdfaoiufoi:        ", cfg.URI)
	_, err := NewClient(&cfg)
	if err != nil {
		t.Error(err)
	}
}
