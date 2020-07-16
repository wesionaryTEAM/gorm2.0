package model

import (
	"gorm.io/gorm"
)

//Model for supplier
type Supplier struct {
	gorm.Model //Provides CreatedAt and UpdatedAt
	Name string `json:"name"`
	Account Account //Supplier has one account
}

type SupplierService interface {
	Create(supplier *Supplier) (*Supplier, error)
	FindAll() ([]Supplier, error)
	Delete(supplier *Supplier) error
}

type SupplierRepository interface {
	Save(supplier *Supplier) (*Supplier, error)
	FindAll() ([]Supplier, error)
	Delete(supplier *Supplier) error
	Migrate() error
}