package config

import (
	"os"

	"github.com/spf13/viper"
)

type (
	//
	Config struct {
		Messages Messages
		Bot      Bot
		Client   Client
	}
	// Bot messages structure
	Messages struct {
		Start           string `mapstructure:"start"`
		Help            string `mapstructure:"help"`
		Categories      string `mapstructure:"categories"`
		AddInitialError string `mapstructure:"add_initial_error"`
		KindError       string `mapstructure:"kind_error"`
		CategoryError   string `mapstructure:"category_error"`
		DelInitialError string `mapstructure:"del_initial_error"`
		NumberError     string `mapstructure:"number_error"`
	}
	// Bot structure
	Bot struct {
		Port   string
		AppURL string
		Token  string
	}
	// Client structure
	Client struct {
		URI        string
		Name       string
		Collection string
	}
)

// Read .yml config
func ReadConfig(path string, cfg *Config) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("main")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.UnmarshalKey("messages", &cfg.Messages)
	return err
}

// Get bot environment variables
func InitBotEnvVars(cfg *Config) {
	cfg.Bot.Port = os.Getenv("PORT")
	cfg.Bot.AppURL = os.Getenv("APP_URL")
	cfg.Bot.Token = os.Getenv("BOT_TOKEN")
}

// Get client environment variables
func InitClientEnvVars(cfg *Config) {
	cfg.Client.URI = os.Getenv("MONGO_URI")
	cfg.Client.Name = os.Getenv("DATABASE_NAME")
	cfg.Client.Collection = os.Getenv("DATABASE_COLLECTION")
}
