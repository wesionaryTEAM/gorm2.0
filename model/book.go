package model

//Model for Book
type Book struct {
	gorm.Model
	ID int64 `json:"id"`
	Title string `json:"title"`
	AuthorID int
	Author Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type BookService interface {
	Create(book *Book) (*Book, error)
	FindAll() ([]Book, error)
	Delete(book *Book) error
}

type BookRepository interface {
	Save(book *Book) (*Book, error)
	FindAll() ([]Book, error)
	Delete(book *Book) error
	Migrate() error
}