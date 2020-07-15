package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm2.0/model"
)
type SupplierController interface {
	GetSuppliers(c *gin.Context)
	AddSupplier(c *gin.Context)
	DeleteSupplier(c *gin.Context)
}

type supplierController struct {
	supplierService model.SupplierService
}

//Constructor Function
func NewSupplierController (service model.SupplierService) SupplierController {
	return &supplierController{
		supplierService: service,
	}
}

//GET
func (s *supplierController) GetSuppliers(c *gin.Context) {
	suppliers, err := s.supplierService.FindAll()
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Error while fetching the suppliers"})
		return
	}
	
	c.JSON(http.StatusOK, suppliers)
}

//POST
func (s *supplierController) AddSupplier(c *gin.Context) {
	var supplier model.Supplier

	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	

	s.supplierService.Create(&supplier)

	c.JSON(http.StatusOK, supplier)
}

//DESTROY
func (s * supplierController) DeleteSupplier(c *gin.Context) {
	var supplier model.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ErrOnDelete := s.supplierService.Delete(&supplier)
	if ErrOnDelete != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Could not delete the supplier"})
		return 
	}

	c.JSON(http.StatusOK, supplier)
}