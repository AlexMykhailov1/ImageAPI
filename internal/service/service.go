package service

import (
	"github.com/AlexMykhailov1/ImageAPI/internal/repository"
)

// Services stores all services
type Services struct {
}

// NewServices returns a pointer to new Services
func NewServices(repos *repository.Repositories) *Services {
	return &Services{}
}
