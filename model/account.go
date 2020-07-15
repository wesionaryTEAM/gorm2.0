package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model //Provides CreatedAt and UpdatedAt
	ID int64
	AccountNumber string
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