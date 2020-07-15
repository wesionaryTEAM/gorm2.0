package service

import (
	"sync"

	"gorm2.0/model"
)

var once sync.Once

type BookService struct {
	bookRepository model.BookRepository
}

var instance *BookService

//Constructor function
func NewBookService(repository model.BookRepository) model.BookService {
	once.Do(func (){
		instance = &BookService{
			bookRepository: repository,
		}
	})
	return instance
}

func (b *BookService) Create(book *model.Book) (*model.Book, error) {
	return b.bookRepository.Save(book)
}

func (b *BookService) FindAll() ([]model.Book, error) {
	return b.bookRepository.FindAll()
}

func (b *BookService) Delete(book *model.Book) error {
	return b.bookRepository.Delete(book)
}

