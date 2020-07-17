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

//Create
func (a *authorService) Create(author *model.Author) (*model.Author, error) {
	return a.authorRepository.Save(author)
} 

//FindAll
func (a *authorService) FindAll() ([]model.Author, error) {
	return a.authorRepository.FindAll()
}

//Delete
func (a *authorService) Delete(author *model.Author) error {
	return a.authorRepository.Delete(author)
}

//FindById
func (a *authorService) FindById(id int64) ([]model.Author, error) {
	return a.authorRepository.FindById(id)
}

//Author count
func (a *authorService) GetTotalNumberOfAuthors() (int64, error) {
	return a.authorRepository.TotalNumberOfAuthors()
}

//Authors name list
func (a *authorService) GetAuthorsNameList() ([]string, error) {
	return a.authorRepository.AllAuthorsNameList()
}