package pubsub

import (
	"context"
	"encoding/json"
	"log"

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

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	simpleQueueType int, // an enum to represent "durable" or "transient"
) (*amqp.Channel, amqp.Queue, error) {
	pubChannel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Could not create new channel: %v", err)
	}
	var durable, autoDelete, exclusive bool

	if simpleQueueType == 1 {
		durable = true
		autoDelete = false
		exclusive = false
	} else {
		durable = false
		autoDelete = true
		exclusive = true
	}
	pubQueue, err := pubChannel.QueueDeclare(
		queueName, durable, autoDelete, exclusive, false, nil,
	)
	if err != nil {
		log.Fatalf("Could not create queue: %v", err)
	}
	err = pubChannel.QueueBind(queueName, key, exchange, false, nil)
	if err != nil {
		log.Fatalf("Could not create new channel: %v", err)
	}
	return pubChannel, pubQueue, nil
}
