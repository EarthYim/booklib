package booklib

import "booklib/app/model"

type booklibService struct {
	booklibStorage
}

type booklibStorage interface {
	AddBook(book model.Book) error
	EditBook(book model.Book) error
	DeleteBook(book model.Book) error
	GetBooks() ([]model.Book, error)
	GetBookById(id int) (model.Book, error)
}

func NewBooklibService(db booklibStorage) *booklibService {
	return &booklibService{db}
}
