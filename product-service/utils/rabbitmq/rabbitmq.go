package rabbitmq

import (
	"fmt"
	"log"

	_config "github.com/wildanie12/go-microservice/product-service/config"

	"github.com/streadway/amqp"
)

// Create to rabbitmq message broker
// using the declared configuration
func Create(config *_config.Config) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", 
		config.RabbitMQ.User,
		config.RabbitMQ.Pass,
		config.RabbitMQ.Host,
		config.RabbitMQ.Pass,
	))
	if err != nil {
		panic("RabbitMQ: connection error")
	}
	ch, err := conn.Channel()
	if err != nil {
		panic("RabbitMQ: failed to open a channel")
	}
	return conn, ch
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}