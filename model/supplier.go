package model

import (
	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model //Provides CreatedAt and UpdatedAt
	ID int64
	Name string
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