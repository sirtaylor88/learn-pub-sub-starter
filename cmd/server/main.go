package main

import (
	"fmt"
	"log"

	"github.com/sirtaylor88/learn-pub-sub-starter/internal/gamelogic"
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

	gamelogic.PrintServerHelp()
	var words []string
	for {
		words = gamelogic.GetInput()
		if len(words) == 0 {
			continue
		}
		switch words[0] {
		case "pause":
			err = pubsub.PublishJSON(
				pubChannel,
				routing.ExchangePerilDirect,
				routing.PauseKey,
				routing.PlayingState{IsPaused: true},
			)
			if err != nil {
				log.Fatalf("Could not publish JSON data: %v", err)
			}
			log.Printf("Pause message sent!")

		case "resume":
			err = pubsub.PublishJSON(
				pubChannel,
				routing.ExchangePerilDirect,
				routing.PauseKey,
				routing.PlayingState{IsPaused: false},
			)
			if err != nil {
				log.Fatalf("Could not publish JSON data: %v", err)
			}
			log.Printf("Resume message sent!")
		case "quit":
			log.Printf("Exiting!")
			return
		default:
			log.Printf("invalid command!")
		}

	}
}
