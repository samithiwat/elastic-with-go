package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"github.com/samithiwat/elastic-with-go/_example_apps/constant"
	"os"
	"time"
)

func main() {
	conn, err := amqp.DialConfig("amqp://guest:guest@localhost:5672/", amqp.Config{
		Vhost: "newbie",
	})
	if err != nil {
		log.Fatal().
			Msg("Failed to connect to RabbitMQ")
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal().
			Msg("Failed to open a channel")
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		constant.ExchangeName,
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal().
			Msg("Failed to declare an exchange")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := os.ReadFile("./inter-summer-courses.json")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while loading course from json")
	}

	if err := ch.PublishWithContext(ctx,
		constant.ExchangeName,
		"cugetreg.scraper.scrape.pending",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	); err != nil {
		log.Fatal().Msg("Failed to publish a message")
	}

	log.Printf("✅ Sent message\n")
}
