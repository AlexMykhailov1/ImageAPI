package delivery

import (
	"fmt"
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/internal/models/image"
	"github.com/AlexMykhailov1/ImageAPI/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
)

// userHandler provides all user interactions logic
type userHandler struct {
	userService service.UserService
	cfg         *config.Config
}

// NewUserHandler returns a pointer to new userHandler
func NewUserHandler(us service.UserService, cfg *config.Config) *userHandler {
	return &userHandler{us, cfg}
}

// InitUserRoutes initializes routing of user handlers
func (us *userHandler) InitUserRoutes(rg *gin.RouterGroup) {
	// General route
	user := rg.Group("/user")

	// Images routes
	user.GET("/download/images/:id", us.DownloadImage) // Get image from the server
}

// DownloadImage takes given image info passes it to the service layer and sends it for a download after
func (us *userHandler) DownloadImage(c *gin.Context) {
	// Take image id as a parameter in path
	idstr := c.Param("id")

	// Convert string to uuid to check if id is valid
	id, err := uuid.Parse(idstr)
	if err != nil {
		log.Printf("Failed to convert string to uuid:%v", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// Take image quality as a query parameter
	quality := c.DefaultQuery("quality", "100")
	if quality != "100" && quality != "75" && quality != "50" && quality != "25" {
		msg := "Wrong quality provided, quality is set to 100%"
		c.JSON(http.StatusBadRequest, gin.H{"msg": msg})
		quality = "100" // Change quality to 100 if quality is not valid
	}

	// Call service layer to process image information and return image name
	imgName, err := us.userService.DownloadImage(c.Request.Context(), id, quality)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// Headers for downloading
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(imgName))
	// Get Content-Type
	ct, err := image.GetImgType(imgName)
	if err != nil {
		msg := fmt.Sprintf("Failed to detect image type:%v", err.Error())
		log.Printf(msg)
		c.Writer.Header().Set("Content-Type", "application/octet-stream")
	}

	c.Writer.Header().Set("Content-Type", ct)

	imgpath := us.cfg.Path.Img + imgName
	c.FileAttachment(imgpath, imgName)
	c.Status(http.StatusOK)
}
