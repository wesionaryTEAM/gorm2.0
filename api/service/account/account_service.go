package service

import (
	"sync"

	"gorm2.0/model"
)

var once sync.Once

type AccountService struct {
	accountRepository model.AccountRepository
}

var instance *AccountService

//Constructor Function
func NewAccountService(repository model.AccountRepository) model.AccountService {
	once.Do(func(){
		instance = &AccountService{
			accountRepository: repository,
		}
	})
	return instance
}

func (a *AccountService) Create(account *model.Account) (*model.Account, error) {
	return a.accountRepository.Save(account)
}

func (a *AccountService) FindAll() ([]model.Account, error) {
	return a.accountRepository.FindAll()
}
	
func (a *AccountService) Delete(account *model.Account) error {
	return a.accountRepository.Delete(account)
}