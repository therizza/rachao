package messaging

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Channel  *amqp.Channel
	Exchange string
}

func (p *RabbitMQ) ensureChannel() error {
	if p.Channel == nil {
		return fmt.Errorf("RabbitMQ channel is nil")
	}
	if p.Channel.IsClosed() {
		return fmt.Errorf("RabbitMQ channel is closed")
	}
	return nil
}

func (p *RabbitMQ) Publish(exchange, routingKey string, body []byte) error {
	if err := p.ensureChannel(); err != nil {
		return err
	}
	if p.Channel == nil {
		return fmt.Errorf("RabbitMQ channel is not initialized")
	}
	err := p.Channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *RabbitMQ) Consumer(handler func(string), exchange string) error {
	msgs, err := c.Channel.Consume(
		"overall",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			handler(string(msg.Body))
		}
	}()

	return nil
}
