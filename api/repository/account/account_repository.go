package repository

import (
	"gorm.io/gorm"
	"gorm2.0/model"
)

type AccountRepository struct {
	DB *gorm.DB
}

//Constructor function
func NewAccountRepository (db *gorm.DB) model.AccountRepository {
	return &AccountRepository {
		DB: db,
	}
}

func (a *AccountRepository) Save(account *model.Account) (*model.Account, error) {
	return account, a.DB.Create(&account).Error
}

func (a *AccountRepository) FindAll() ([]model.Account, error) {
	var accounts []model.Account
	err := a.DB.Find(&accounts).Error
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (a *AccountRepository) Delete(account *model.Account) error {
	return a.DB.Delete(&account).Error
}

func (a *AccountRepository) Migrate() error {
	err := a.DB.AutoMigrate(&model.Account{})
	return err
}