package delivery

import (
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// uploadHandler provides all file upload delivery logic
type uploadHandler struct {
	uploadService service.UploadService
	cfg           *config.Config
}

// NewUploadHandler returns a pointer to new UploadHandler
func NewUploadHandler(us service.UploadService, cfg *config.Config) *uploadHandler {
	return &uploadHandler{uploadService: us, cfg: cfg}
}

// InitUploadRoutes initializes routing of upload handlers
func (us *uploadHandler) InitUploadRoutes(rg *gin.RouterGroup) {
	// General route
	upload := rg.Group("/upload")

	// Images routes
	upload.POST("/images", us.UploadImage) // Upload one image
}

// UploadImage handles image uploading to the server and passes it to the next layer for further processing
func (us *uploadHandler) UploadImage(c *gin.Context) {
	// Set maximum image size
	const maxbytes int64 = 1024 * 1024 * 8 // 8 Mb
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxbytes)

	// Upload image
	file, err := c.FormFile("image")
	if err != nil {
		log.Printf("Could not upload image: %v\n", err.Error())
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})

	}

	// Get Content-Type
	ct := file.Header.Get("Content-Type")

	// Check whether it is a picture
	if ct != "image/png" && ct != "image/jpeg" {
		msg := "Uploaded file is not a picture"
		log.Printf(msg)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": msg})
		return
	}

	// Pass image to service layer
	id, err := us.uploadService.UploadImage(c, file)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Picture uploaded successfully!", "id": id})
}
