package infrastructure

import (
	"net/http"
	"os"

	// "gorm2.0/api/routes"
	router "gorm2.0/http"
	"gorm2.0/api/routes"
	"github.com/gin-gonic/gin"
  "gorm.io/gorm"
)


//Setup routes
func SetupRoutes(db *gorm.DB) {
	httpRouter := router.NewGinRouter()

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})

	// Author Routers
	routes.AuthorRoutes(httpRouter.GROUP("/authors"),db)

	// Book routes
	routes.BookRoutes(httpRouter.GROUP("/books"), db)

	//Supplier Routes
	routes.SupplierRoutes(httpRouter.GROUP("/suppliers"), db)

	// Start the server
	port := os.Getenv("SERVER_PORT")
	httpRouter.SERVE(port)
	if port == "" {
		httpRouter.SERVE(":8000")
	} else {
		httpRouter.SERVE(port)
	}
}