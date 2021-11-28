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
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var cfg *Config
	err = viper.UnmarshalKey("messages", &cfg.Messages)
	if err != nil {
		return nil, err
	}
	cfg.Bot.Port = viper.Get("PORT").(string)
	cfg.Bot.AppURL = viper.Get("APP_URL").(string)
	cfg.Bot.Token = viper.Get("BOT_TOKEN").(string)

	cfg.Client.URI = viper.Get("MONGO_URI").(string)
	cfg.Client.Name = viper.Get("DATABASE_NAME").(string)
	cfg.Client.Collection = viper.Get("DATABASE_COLLECTION").(string)
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
