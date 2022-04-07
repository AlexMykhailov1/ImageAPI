package delivery

import (
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/internal/service"
	"github.com/gin-gonic/gin"
)

// UploadHandler stores all methods of uploadHandler struct
type UploadHandler interface {
	InitUploadRoutes(rg *gin.RouterGroup)
	UploadImage(c *gin.Context)
}

// Handlers stores all handler interfaces
type Handlers struct {
	UploadHandler UploadHandler
}

// NewHandlers returns a pointer to new Handlers
func NewHandlers(services *service.Services, cfg *config.Config) *Handlers {
	return &Handlers{UploadHandler: NewUploadHandler(services.UploadService, cfg)}
}

// InitRoutes initializes the routing of the application
func (h *Handlers) InitRoutes(rg *gin.RouterGroup) {
	h.UploadHandler.InitUploadRoutes(rg)
}
