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

//Save
func (a *AuthorRepository) Save(author *model.Author) (*model.Author, error) {
	return author, a.DB.Create(author).Error
}

//FindAll
func (a *AuthorRepository) FindAll() ([]model.Author, error) {
	var authors []model.Author
	err := a.DB.Find(&authors).Error
	return authors, err
}

//Delete
func (a *AuthorRepository) 	Delete (author *model.Author) error {
	return a.DB.Delete(&author).Error
}

//FindById
func (a *AuthorRepository) FindById(id int64) ([]model.Author, error) {
	var author []model.Author
	err := a.DB.Where("id = ?", id).Find(&author).Error
	// if err != nil {
	// 	return nil, err
	// }
	return author, err
}

//Migrate
func (a *AuthorRepository) Migrate() error {
	err := a.DB.AutoMigrate(&model.Author{})
	return err
}