package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model //Provides CreatedAt and UpdatedAt
	AccountNumber string `gorm:"unique;"`
	SupplierID int
}

type AccountService interface {
	Create(account *Account) (*Account, error)
	FindAll() ([]Account, error)
	Delete(account *Account) error
}

type AccountRepository interface {
	Save(account *Account) (*Account, error)
	FindAll() ([]Account, error)
	Delete(account *Account) error 
	Migrate() error
}