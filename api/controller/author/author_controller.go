package controller

import (
	"net/http"

	"gorm2.0/model"
	"github.com/gin-gonic/gin"
)

type AuthorController interface {
	GetAuthors(c *gin.Context)
	AddAuthor(c *gin.Context)
	DeleteAuthor(c *gin.Context)
}

type authorController struct {
	authorService model.AuthorService
}

//Constructor Function
func NewAuthorController(service model.AuthorService) AuthorController {
	return &authorController{
		authorService: service,
	}
}

//GET
func (a *authorController) GetAuthors(c *gin.Context) {
	authors, err := a.authorService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while fetching the authors"})
		return 
	}

	c.JSON(http.StatusOK, authors)
}

//POST
func (a *authorController)	AddAuthor(c *gin.Context) {
	var author model.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	a.authorService.Create(&author)

	c.JSON(http.StatusOK, author)
}

//DESTORY
func (a *authorController) DeleteAuthor(c *gin.Context) {
	var author model.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	ErrOnDelete := a.authorService.Delete(&author)
	if ErrOnDelete != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not delete the author"})
		return 
	}
	c.JSON(http.StatusOK, author)
}
