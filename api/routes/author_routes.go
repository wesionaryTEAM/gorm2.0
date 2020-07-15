package routes

import (
	"log"

	author_repository "gorm2.0/api/repository/author"
	author_controller "gorm2.0/api/controller/author"
	author_service "gorm2.0/api/service/author"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthorRoutes(route *gin.RouterGroup, db *gorm.DB) {
	//dependency injection 
	authorRepository := author_repository.NewAuthorRepository(db)
	if err := authorRepository.Migrate(); err != nil {
		log.Fatal("Author migration failed", err)
	}

	authorService := author_service.NewAuthorService(authorRepository)
	authorController := author_controller.NewAuthorController(authorService)

	//Routes
	route.GET("/", authorController.GetAuthors)
	route.POST("/", authorController.AddAuthor)
	route.DELETE("/", authorController.DeleteAuthor)
}