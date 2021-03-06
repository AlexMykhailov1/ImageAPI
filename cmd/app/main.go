package main

import (
	"database/sql"
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/internal/rabbit"
	"github.com/AlexMykhailov1/ImageAPI/internal/server"
	"github.com/AlexMykhailov1/ImageAPI/storage/postgres"
	"github.com/gin-gonic/gin"
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

	// Connect to postgres
	db, err := postgres.NewPostgres(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err.Error())
	}

	// Defer close database connection
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v\n", err.Error())
		}
	}(db)

	// Create new server instance
	s := server.NewServer(gin.Default(), db, cfg, rb)

	// Run server
	if err = s.Start(); err != nil {
		log.Fatalf("Error starting server: %v\n", err.Error())
	}
}
