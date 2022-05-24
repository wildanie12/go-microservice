package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		BaseURL string
		Port    string
		URL     string
	}
	Mongo struct {
		Host string
		Port string
		User string
		Pass string
		DB   string
	}
}
var config *Config

func New() *Config {
	if config == nil {
		config = initConfig()
	}
	return config
}

func initConfig() *Config {
	config := Config{}

	err := godotenv.Load()
	if err != nil {
		config.App.BaseURL = "localhost"
		config.App.Port = "8000"
		config.App.URL = "localhost:8000"

		config.Mongo.Host = "localhost:8000"
		config.Mongo.Port = "8000"
		config.Mongo.User = "root"
		config.Mongo.Pass = ""
		config.Mongo.DB = "product"
	}
	config.App.BaseURL = os.Getenv("APP_BASE_URL")
	config.App.Port = os.Getenv("APP_PORT")
	config.App.URL = os.Getenv("APP_URL")
	config.Mongo.Host = os.Getenv("MONGO_HOST")
	config.Mongo.Port = os.Getenv("MONGO_PORT")
	config.Mongo.User = os.Getenv("MONGO_USER")
	config.Mongo.Pass = os.Getenv("MONGO_PASS")
	config.Mongo.DB = os.Getenv("MONGO_DB")

	return &config
}