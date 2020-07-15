package routes


import (
	"log"

	book_repository "gorm2.0/api/repository/book"
	book_controller "gorm2.0/api/controller/book"
	book_service "gorm2.0/api/service/book"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookRoutes(route *gin.RouterGroup, db *gorm.DB) {
	//dependency injection 
	bookRepository := book_repository.NewBookRepository(db)
	if err := bookRepository.Migrate(); err != nil {
		log.Fatal("Book migration failed", err)
	}

	bookService := book_service.NewBookService(bookRepository)
	bookController := book_controller.NewBookController(bookService)

	//Routes
	route.GET("/", bookController.GetBooks)
	route.POST("/", bookController.AddBook)
	route.DELETE("/", bookController.DeleteBook)
}