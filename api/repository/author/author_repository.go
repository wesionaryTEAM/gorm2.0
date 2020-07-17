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
	return author, err
}

//Count the author's total number
func (a *AuthorRepository) TotalNumberOfAuthors() (int64, error) {
	var count int64
	err := a.DB.Table("authors").Count(&count).Error
	// err := a.DB.Model(&Authors{}).Count(&count).Error
	return count, err
}

//Pluck author's names only
func (a *AuthorRepository) AllAuthorsNameList() ([]string, error) {
	var names []string
	err := a.DB.Table("authors").Pluck("name", &names).Error
	// err := a.DB.Model(&Authors{}).Pluck("name", &names).Error
	return names, err
}

//Migrate
func (a *AuthorRepository) Migrate() error {
	err := a.DB.AutoMigrate(&model.Author{})
	return err
}