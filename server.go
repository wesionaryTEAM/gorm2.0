package main

import (
	"gorm2.0/infrastructure"
	"gorm2.0/utils"
)

func main() {
	utils.LoadEnv()
	
	db := infrastructure.GetDatabaseInstance()

	infrastructure.SetupRoutes(db)
}