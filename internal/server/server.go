package server

import (
	"database/sql"
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/internal/delivery"
	"github.com/AlexMykhailov1/ImageAPI/internal/rabbit"
	"github.com/AlexMykhailov1/ImageAPI/internal/repository"
	"github.com/AlexMykhailov1/ImageAPI/internal/service"
	"github.com/gin-gonic/gin"
)

type server struct {
	*gin.Engine
	db  *sql.DB
	cfg *config.Config
	rb  *rabbit.Rabbit
}

// NewServer returns a pointer to new Server
func NewServer(router *gin.Engine, db *sql.DB, cfg *config.Config, rb *rabbit.Rabbit) *server {
	return &server{router, db, cfg, rb}
}

// Start maps handlers, creates needed RabbitMQ queues and calls Run() to start the server
func (s *server) Start() error {
	s.mapHandlers(s.Engine)
	// Set up all rabbit queues
	err := s.rb.SetUpQueues(s.cfg)
	if err != nil {
		return err
	}
	// Run the server
	if err = s.Run(s.cfg.SRV.Port); err != nil {
		return err
	}
	return nil
}

// mapHandlers initializes all routes using InitRoutes()
func (s *server) mapHandlers(router *gin.Engine) {
	// Init layers
	repos := repository.NewRepositories(s.db)
	services := service.NewServices(repos, s.cfg, s.rb)
	handlers := delivery.NewHandlers(services, s.cfg)

	// Initial path
	v1 := router.Group("api/v1")

	// Init all routes
	handlers.InitRoutes(v1)
}
