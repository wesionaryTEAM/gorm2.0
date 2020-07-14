package repository

import (
	"gorm2.0/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

//Constructor function
func NewBookRepository (db *gorm.DB) model.BookRepository {
	return &BookRepository {
		DB: db,
	}
}

func (b *BookRepository) Save(book *Book) (*Book, error) {
	return book, b.DB.Create(&book).Error
}

func (b *BookRepository) FindAll() ([]Book, error) {
	var books = model.Book
	err := b.DB.Find(&books).Error
}

func (b *BookRepository) Delete(book *Book) error {
	return b.DB.Delete(&book).Error
}

func (b *BookRepository) 	Migrate() error {
	return b.DB.AutoMigrate(&model.User{}).Error
}