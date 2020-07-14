package service

import (
	"sync"

	"gorm2.0/model"
)

var once sync.once

type BookService struct {
	bookRepository model.BookRepository
}

var instance *BookService

//Constructor function
func NewBookService(repository model.BookRepository) model.BookService {
	once.Do(func (){
		isntance = &BookService{
			bookRepository: repository,
		}
	})
	return instance
}

func (b *BookService) Create(book *Book) (*Book, error) {
	b.bookRepository.Save(book)
}

func (b *BookService) FindAll() ([]Book, error) {
	return b.bookRepository.FindAll()
}

func (b *BookService) Delete(book *Book) error {
	return b.bookRepository.Delete(book)
}

