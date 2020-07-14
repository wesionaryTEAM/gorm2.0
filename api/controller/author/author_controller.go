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

type AuthorController struct {
	authorService model.AuthorService
}

//Constructor Function
func NewAuthorController(service model.AuthorService) AuthorController {
	return &AuthorController{
		authorService: service,
	}
}

//GET
func (a *AuthorController) GetAuthors(c *gin.Context) {
	authors, err := a.authorService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while fetching the authors"})
		return 
	}

	c.JSON(http.StatusOK, authors)
}

//POST
func (a *AuthorController)	AddAuthor(c *gin.Context) {
	var author model.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	author, err := a.authorService.Create(&author)

	c.JSON(http.StatusOK, author)
}

//DESTORY
func (a *AuthorController) DeleteAuthor(c *gin.Context) {
	var author model.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	ErrOnDelete := a.authorService.Delete(author)
	if ErrOnDelete != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not delete the author"})
		return 
	}
	c.JSON(http.StatusOK, author)
}
