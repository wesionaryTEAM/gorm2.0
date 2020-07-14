package repository

import (
	"gorm2.0/model"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	DB *gorm.DB
}

// Constructor function
func NewAuthorRepository (db *gorm.DB) model.AuthorRepository {
	return &AuthorRepository {
		DB: db,
	}
}

func (a *AuthorRepository) Save(author *Author) (*Author, error) {
	return author, a.DB.Create(author).Error
}

func (a *AuthorRepository) FindAll() ([]Author, error) {
	var authors []model.Author
	err := a.DB.Find(&authors).Error
	return authors, err
}

func (a *AuthorRepository) 	Delete (author *Author) error {
	return a.DB.Delete(&author).Error
}

func (a *AuthorRepository) Migrate() error {
	return a.DB.AutoMigrate(&model.User{}).Error
}