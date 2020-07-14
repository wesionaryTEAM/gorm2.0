package infrastructure

import (
	"net/http"
	"os"

	router "gorm2.0/http"
	"github.com/gin-gonic/gin"
  "gorm.io/gorm"
)


//Setup routes
func SetupRoutes(db *gorm.DB) {
	httpRouter := router.NewGinRouter()

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})

	// Start the server
	port := os.Getenv("SERVER_PORT")
	httpRouter.SERVE(port)
	if port == "" {
		httpRouter.SERVE(":8000")
	} else {
		httpRouter.SERVE(port)
	}
}