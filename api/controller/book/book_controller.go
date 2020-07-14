package controller

import (
	"net/http"

	"gorm2.0/model"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	GetBooks(c *gin.Context)
	AddBook(c *gin.Context)
	DeleteBook(c *gin.Context)
}

type BookController struct {
	bookService model.bookService
}

//Constructor Function
func NewBookController(service bookService) BookController {
	return &BookController{
		bookService: service,
	}
}

//GET
func (b *BookController) GetBooks(c *gin.Context) {
	books, err := b.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while fetching the books"})
		return 
	}

	c.JSON(http.StatusOK, books)
}

//POST 
func (b *BookController) AddBook(c * gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := b.bookService.Create(&book)

	c.JSON(http.StatusOk, book)
}

//DESTROY
func (b *BookController) DeleteBook(c *gin.Context){
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ErrOnDelete := b.bookService.Delete(book)
	if ErrOnDelete != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not delete the book"})
		return 
	}

	c.JSON(http.StatusOK, book)
}