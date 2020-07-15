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

func (b *BookRepository) Save(book *model.Book) (*model.Book, error) {
	return book, b.DB.Create(&book).Error
}

func (b *BookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	err := b.DB.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b *BookRepository) Delete(book *model.Book) error {
	err := b.DB.Delete(&book).Error
	return err
}

func (b *BookRepository) 	Migrate() error {
	err := b.DB.AutoMigrate(&model.Book{})
	return err
}