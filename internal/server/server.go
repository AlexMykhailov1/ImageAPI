package server

import (
	"database/sql"
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/internal/delivery"
	"github.com/AlexMykhailov1/ImageAPI/internal/repository"
	"github.com/AlexMykhailov1/ImageAPI/internal/service"
	"github.com/gin-gonic/gin"
)

type server struct {
	*gin.Engine
	db  *sql.DB
	cfg *config.Config
}

// NewServer returns a pointer to new Server
func NewServer(router *gin.Engine, db *sql.DB, cfg *config.Config) *server {
	return &server{router, db, cfg}
}

// Start maps handlers and calls Run() to start the server
func (s *server) Start() error {
	s.mapHandlers(s.Engine)
	if err := s.Run(s.cfg.SRV.Port); err != nil {
		return err
	}
	return nil
}

// mapHandlers initializes all routes using InitRoutes()
func (s *server) mapHandlers(router *gin.Engine) {
	// Init layers
	repos := repository.NewRepositories(s.db)
	services := service.NewServices(repos, s.cfg)
	handlers := delivery.NewHandlers(services, s.cfg)

	// Initial path
	v1 := router.Group("api/v1")

	// Init all routes
	handlers.InitRoutes(v1)
}
