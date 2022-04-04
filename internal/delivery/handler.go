package delivery

import (
	"github.com/AlexMykhailov1/ImageAPI/internal/service"
	"github.com/gin-gonic/gin"
)

// Handlers stores all handlers
type Handlers struct {
}

// NewHandlers returns a pointer to new Handlers
func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{}
}

// InitRoutes initializes the routing of the application
func (h *Handlers) InitRoutes(g *gin.RouterGroup) {

}
