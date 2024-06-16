package storage

import (
	"booklib/app/model"

	"gorm.io/gorm"
)

type booklibDb struct {
	*gorm.DB
}

func NewBooklibDb(db *gorm.DB) *booklibDb {
	return &booklibDb{db}
}

func (db *booklibDb) AddBook(book model.Book) error {
	return db.Create(&book).Error
}

func (db *booklibDb) EditBook(book model.Book) error {
	return db.Updates(&book).Error
}

func (db *booklibDb) DeleteBook(book model.Book) error {
	return db.Delete(&book).Error
}

func (db *booklibDb) GetBooks() ([]model.Book, error) {
	var books []model.Book
	err := db.Find(&books).Error
	return books, err
}

func (db *booklibDb) GetBookById(id int) (model.Book, error) {
	var book model.Book
	err := db.Find(&book, id).Error
	return book, err
}

// func (db *booklibDb) GetBookByTitle(title string) (model.Book, error) {
// 	var book model.Book
// 	err := db.Where("title = ?", title).Find(&book).Error
// 	return book, err
// }

// func (db *booklibDb) GetBookByYear(year string) ([]model.Book, error) {
// 	var books []model.Book
// 	err := db.Where("year = ?", year).Find(&books).Error
// 	return books, err
// }
