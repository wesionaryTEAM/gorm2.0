package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	author_cache "gorm2.0/api/cache"
	author_controller "gorm2.0/api/controller/author"
	author_repository "gorm2.0/api/repository/author"
	author_service "gorm2.0/api/service/author"
)

// Author Routes
func AuthorRoutes(route *gin.RouterGroup, db *gorm.DB) {
	//dependency injection
	authorRepository := author_repository.NewAuthorRepository(db)
	if err := authorRepository.Migrate(); err != nil {
		log.Fatal("Author migration failed", err)
	}

	authorService := author_service.NewAuthorService(authorRepository)
	authorCache := author_cache.NewRedisCache("localhost:6379", 1, 10)
	authorController := author_controller.NewAuthorController(authorService, authorCache)

	//Routes
	route.GET("/", authorController.GetAuthors)
	route.GET("/getTotalAuthorCount", authorController.GetTotalNumberOfAuthors)
	route.GET("/getAuthorsNameList", authorController.GetAuthorsNameList)
	route.POST("/", authorController.AddAuthor)
	route.POST("/findById", authorController.FindById)
	route.DELETE("/", authorController.DeleteAuthor)
}
