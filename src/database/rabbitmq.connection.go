package database

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/samithiwat/elastic-with-go/src/config"
)

func InitRabbitMQConnection() (*amqp.Connection, error) {
	conf, err := config.LoadRabbitMQConfig()
	if err != nil {
		return nil, err
	}

	return amqp.DialConfig(conf.Host, amqp.Config{
		Vhost: conf.VHost,
	})
}
