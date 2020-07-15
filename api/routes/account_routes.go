package routes

import (
	"log"

	account_repository "gorm2.0/api/repository/account"
	account_controller "gorm2.0/api/controller/account"
	account_service "gorm2.0/api/service/account"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AccountRoutes(route *gin.RouterGroup, db *gorm.DB) {
	//dependency injection setup
	accountRepository := account_repository.NewAccountRepository(db)
	if err := accountRepository.Migrate(); err != nil {
		log.Fatal("Account Migration Failed", err)
	}

	accountService := account_service.NewAccountService(accountRepository)
	accountController := account_controller.NewAccountController(accountService)

	//Routes
	route.GET("/", accountController.GetAccounts)
	route.POST("/", accountController.AddAccount)
	route.DELETE("/", accountController.DeleteAccount)
}