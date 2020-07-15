package service

import (
	"sync"

	"gorm2.0/model"
)

var once sync.Once  

type authorService struct {
	authorRepository model.AuthorRepository
}

var instance *authorService

//Constructor function
func NewAuthorService(repository model.AuthorRepository) model.AuthorService {
	once.Do(func() {
		instance = &authorService{
			authorRepository: repository,
		}
	})
	return instance
}

func (a *authorService) Create(author *model.Author) (*model.Author, error) {
	return a.authorRepository.Save(author)
} 

func (a *authorService) FindAll() ([]model.Author, error) {
	return a.authorRepository.FindAll()
}

func (a *authorService) Delete(author *model.Author) error {
	return a.authorRepository.Delete(author)
}
