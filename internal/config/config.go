package config

import (
	"os"

	"github.com/spf13/viper"
)

func InitConfig() (*Config, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	viper.AddConfigPath(path + "/configs")
	viper.SetConfigName("main")
	viper.SetConfigType("yml")
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var cfg *Config
	err = viper.UnmarshalKey("messages", &cfg.Messages)
	if err != nil {
		return nil, err
	}
	cfg.Bot.Port = os.Getenv("PORT")
	cfg.Bot.AppURL = os.Getenv("APP_URL")
	cfg.Bot.Token = os.Getenv("BOT_TOKEN")
	cfg.Client.URI = os.Getenv("MONGO_URI")
	cfg.Client.Name = os.Getenv("DATABASE_NAME")
	cfg.Client.Collection = os.Getenv("DATABASE_COLLECTION")
	return cfg, nil
}

type (
	Config struct {
		Messages Messages
		Bot      Bot
		Client   Client
	}
	Messages struct {
		Start         string `mapstructure:"start"`
		Help          string `mapstructure:"help"`
		Categories    string `mapstructure:"categories"`
		InitialError  string `mapstructure:"initial_error"`
		CategoryError string `mapstructure:"category_error"`
	}
	Bot struct {
		Port   string
		AppURL string
		Token  string
	}
	Client struct {
		URI        string
		Name       string
		Collection string
	}
)
