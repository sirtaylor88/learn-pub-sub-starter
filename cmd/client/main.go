package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/sirtaylor88/learn-pub-sub-starter/internal/gamelogic"
	"github.com/sirtaylor88/learn-pub-sub-starter/internal/pubsub"
	"github.com/sirtaylor88/learn-pub-sub-starter/internal/routing"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril client...")
	const rabbitConnString = "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(rabbitConnString)
	if err != nil {
		log.Fatalf("Could not connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
	fmt.Println("Peril game server connected successfully to RabbitMQ !!!")

	username, err := gamelogic.ClientWelcome()
	if err != nil {
		log.Fatalf("Invalid username: %v", err)
	}

	_, _, err = pubsub.DeclareAndBind(
		conn,
		routing.ExchangePerilDirect,
		routing.PauseKey+"."+username,
		routing.PauseKey,
		0,
	)

	if err != nil {
		log.Printf("Could not declare and bind the queue to exchange: %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Connection is shutting down...")
}
