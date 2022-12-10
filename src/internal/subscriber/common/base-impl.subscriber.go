package common

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type MessageHandler func(context.Context)

type subscriber struct {
	ch           *amqp.Channel
	errCh        chan error
	topicList    []string
	exchangeKind string
	exchangeName string
	queueName    string
	handlerList  []MessageHandler
}

func NewSubscriber(connection *amqp.Connection, errCh chan error, topicList []string, exchangeKind string, exchangeName string) (Subscriber, error) {
	ch, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return &subscriber{
		ch:           ch,
		errCh:        errCh,
		topicList:    topicList,
		exchangeKind: exchangeKind,
		exchangeName: exchangeName,
	}, nil
}

func (s *subscriber) Listen() {
	msgs, err := s.initConsumer()
	if err != nil {
		s.errCh <- err
	}

	log.Info().
		Str("queue_name", s.queueName).
		Str("exchange", s.exchangeName).
		Msg("Ready to accept message from RabbitMQ")

	for d := range msgs {
		log.Info().
			Str("topic", d.RoutingKey).
			Msg("Received message")

		for _, messageHandler := range s.handlerList {
			messageHandler(context.WithValue(context.Background(), "message", d.Body))
		}
	}
}

func (s *subscriber) Close() error {
	return s.ch.Close()
}

func (s *subscriber) initConsumer() (<-chan amqp.Delivery, error) {
	if err := s.ch.ExchangeDeclare(
		s.exchangeName,
		s.exchangeKind,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		log.Error().
			Err(err).
			Str("exchange", s.exchangeName).
			Msg("Failed to create exchange")
		return nil, err
	}

	queue, err := s.ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to create queue")
		return nil, err
	}

	s.queueName = queue.Name

	for _, topic := range s.topicList {
		log.Info().
			Str("queue_name", queue.Name).
			Str("exchange", s.exchangeName).
			Str("topic", topic).
			Msg("Binding queue")

		if err := s.ch.QueueBind(
			queue.Name,
			topic,
			s.exchangeName,
			false,
			nil,
		); err != nil {
			log.Error().
				Err(err).
				Str("queue_name", queue.Name).
				Str("exchange", s.exchangeName).
				Str("topic", topic).
				Msg("Failed to binding queue")
		}
	}

	return s.ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
}

func (s *subscriber) RegisterHandler(handler ...MessageHandler) {
	for _, messageHandler := range handler {
		s.handlerList = append(s.handlerList, messageHandler)
	}
}
