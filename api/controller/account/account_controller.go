package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm2.0/model"
)

type AccountController interface {
	GetAccounts(c *gin.Context)
	AddAccount(c *gin.Context)
	DeleteAccount(c *gin.Context)
}

type accountController struct {
	accountService model.AccountService
}

//Constructor Function
func NewAccountController(service model.AccountService) AccountController {
	return &accountController{
		accountService: service,
	}
}

//GET
func (a *accountController) GetAccounts(c *gin.Context) {
	accounts, err := a.accountService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while fetching the accounts"})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

//POST
func (a *accountController) AddAccount(c *gin.Context) {
	var account model.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.accountService.Create(&account)

	c.JSON(http.StatusOK, account)
}

//DELETE
func (a *accountController) DeleteAccount(c *gin.Context) {
	var account model.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 

	ErrOnDelete := a.accountService.Delete(&account)
	if ErrOnDelete!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not delete the account"})
		return
	}

	c.JSON(http.StatusOK, account)
}