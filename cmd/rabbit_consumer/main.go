package main

import (
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/internal/rabbit"
	"github.com/streadway/amqp"
	"log"
)

const envFilePath = "./env/"

func main() {
	// Initialize config file
	cfg, err := config.LoadConfig(envFilePath)
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err.Error())
	}

	// Connect to RabbitMQ
	rb, err := rabbit.NewRabbit(cfg)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v\n", err.Error())
	}

	// Defer close RabbitMQ connection
	defer func(conn *amqp.Connection) {
		err = conn.Close()
		if err != nil {
			log.Printf("Error closing RabbitMQ connection: %v\n", err.Error())
		}
	}(rb.Connection)

	// Start consume messages from image queue
	err = rb.ConsumeImgID(cfg)
	if err != nil {
		log.Fatalf("Error consuming messages:%v", err.Error())
	}
}
