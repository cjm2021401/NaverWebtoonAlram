package config

import (
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Database Database
	NaverWebtoon NaverWebtoon
	SlackWebhook SlackWebhook
	LineToken LineToken
}
type Database struct {
	Url      string
	User     string
	Password string
	Name     string
}

type NaverWebtoon struct {
	Url string
}
type SlackWebhook struct {
	Url string
}
type LineToken struct {
	Token string
}

var DB *pg.DB
var Env Config

func GetEnvironmentVariable() error {
	viper.SetConfigType("json")
	viper.AddConfigPath("config")
	viper.SetConfigName("dev.json")
	var config Config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	Env = config

	return nil
}
