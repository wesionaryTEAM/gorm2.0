package model

import (
	"gorm.io/gorm"
)
//Model for Author
type Author struct {
	gorm.Model
	ID int64 `json:"id"`
	Name	string `json:"name"`
}

type AuthorService interface {
	Create(author *Author) (*Author, error)
	FindAll() ([]Author, error)
	Delete(author *Author) error
}

type AuthorRepository interface {
	Save(author *Author) (*Author, error)
	FindAll() ([]Author, error)
	Delete (author *Author) error
	Migrate() error
}