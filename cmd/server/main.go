package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/sirtaylor88/learn-pub-sub-starter/internal/pubsub"
	"github.com/sirtaylor88/learn-pub-sub-starter/internal/routing"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")
	const rabbitConnString = "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(rabbitConnString)
	if err != nil {
		log.Fatalf("Could not connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
	fmt.Println("Peril game server connected successfully to RabbitMQ !!!")

	pubChannel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Could not create new channel: %v", err)
	}
	err = pubsub.PublishJSON(
		pubChannel,
		routing.ExchangePerilDirect,
		routing.PauseKey,
		routing.PlayingState{IsPaused: true},
	)

	if err != nil {
		log.Fatalf("Could not publish JSON data: %v", err)
	}
	fmt.Println("Pause message sent!")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Connection is shutting down...")
}
