package pubsub

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	data, err := json.Marshal(val)
	if err != nil {
		return err
	}
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        data,
	}

	ctx := context.Background()

	return ch.PublishWithContext(ctx, exchange, key, false, false, msg)
}
