package main

import (
	"database/sql"
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/internal/server"
	"github.com/AlexMykhailov1/ImageAPI/storage/postgres"
	"github.com/gin-gonic/gin"
	"log"
)

const envFilePath = "./env/"

func main() {
	// Initialize config file
	cfg, err := config.LoadConfig(envFilePath)
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	// New postgres connection
	db, err := postgres.NewPostgres(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	// Defer close database connection
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Printf("Error closing database: %v\n", err)
		}
	}(db)

	// Create new server instance
	s := server.NewServer(gin.Default(), db, cfg)

	// Run server
	if err = s.Start(); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
