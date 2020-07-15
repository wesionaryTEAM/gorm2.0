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

func (a *AuthorRepository) Save(author *model.Author) (*model.Author, error) {
	return author, a.DB.Create(author).Error
}

func (a *AuthorRepository) FindAll() ([]model.Author, error) {
	var authors []model.Author
	err := a.DB.Find(&authors).Error
	return authors, err
}

func (a *AuthorRepository) 	Delete (author *model.Author) error {
	return a.DB.Delete(&author).Error
}

func (a *AuthorRepository) Migrate() error {
	err := a.DB.AutoMigrate(&model.Author{})
	return err
}