package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config main struct to store
// configuration based on env file
type Config struct {
	App struct {
		BaseURL string
		Port    string
		URL     string
		Env		string
	}
	Mongo struct {
		Host string
		Port string
		User string
		Pass string
		DB   string
	}
	RabbitMQ struct {
		User string
		Pass string
		Host string
		Port string
	}
}
var config *Config

// New returns an instance of main config object
// It will return existing config if it's already initiated
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
		config.App.Env = "production"

		config.Mongo.Host = "localhost:8000"
		config.Mongo.Port = "8000"
		config.Mongo.User = "root"
		config.Mongo.Pass = ""
		config.Mongo.DB = "product"

		config.RabbitMQ.User = "user"
		config.RabbitMQ.Pass = "user"
		config.RabbitMQ.Host = "localhost"
		config.RabbitMQ.Port = "5672"
	}
	config.App.BaseURL = os.Getenv("APP_BASE_URL")
	config.App.Port = os.Getenv("APP_PORT")
	config.App.URL = os.Getenv("APP_URL")
	config.App.Env = os.Getenv("APP_ENV")
	config.Mongo.Host = os.Getenv("MONGO_HOST")
	config.Mongo.Port = os.Getenv("MONGO_PORT")
	config.Mongo.User = os.Getenv("MONGO_USER")
	config.Mongo.Pass = os.Getenv("MONGO_PASS")
	config.Mongo.DB = os.Getenv("MONGO_DB")
	config.RabbitMQ.User = os.Getenv("RABBITMQ_USER")
	config.RabbitMQ.Pass = os.Getenv("RABBITMQ_PASS")
	config.RabbitMQ.Host = os.Getenv("RABBITMQ_HOST")
	config.RabbitMQ.Port = os.Getenv("RABBITMQ_PORT")

	return &config
}