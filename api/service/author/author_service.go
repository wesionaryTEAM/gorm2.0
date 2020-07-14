package service

import (
	"sync"

	"gorm2.0/model"
)

var once sync.Once  

type AuthorService struct {
	authorRepository model.AuthorRepository
}

var instance *AuthorService

//Constructor function
func NewAuthorService(repository model.AuthorService) model.AuthorService {
	once.Do(func() {
		instance = &AuthorService{
			authorRepository: repository,
		}
	})
	return instance
}

func (a *AuthorService) Create(author *model.Author) (*model.Author, error) {
	return a.authorRepository.Save(author)
} 

func (a *AuthorService) FindAll() ([]Author, error) {
	return a.authorRepository.FindAll()
}

func (a *AuthorService) Delete(author *Author) error {
	return a.authorRepository.Delete(author)
}
