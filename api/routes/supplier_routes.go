package routes

import (
	"log"

	supplier_repository "gorm2.0/api/repository/supplier"
	supplier_service "gorm2.0/api/service/supplier"
	supplier_controller "gorm2.0/api/controller/supplier"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SupplierRoutes(route *gin.RouterGroup, db *gorm.DB){
	//dependency injection setup part
	supplierRepository := supplier_repository.NewSupplierRepository(db)
	if err := supplierRepository.Migrate(); err != nil {
		log.Fatal("Supplier migration failed", err)
	}

	supplierService := supplier_service.NewSupplierService(supplierRepository)
	supplierController := supplier_controller.NewSupplierController(supplierService)

	//Routes
	route.GET("/", supplierController.GetSuppliers)
	route.POST("/", supplierController.AddSupplier)
	route.DELETE("/", supplierController.DeleteSupplier)
}