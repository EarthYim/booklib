package booklib

import (
	"booklib/app/model"
	"booklib/app/model/response"
)

func (s *booklibService) Get() []response.Get {
	return []response.Get{}
}

func (s *booklibService) GetByID(id int) (*response.Get, error) {
	book, err := s.GetBookById(id)
	if err != nil {
		return nil, err
	}
	resp := &response.Get{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Genre:  book.Genre,
		Year:   book.Year,
		Tag:    book.Tag,
	}
	return resp, nil
}

// func (s *booklibService) GetByTag(tag string) {}

func (s *booklibService) GetAll() (*[]response.Get, error) {

	books, err := s.GetBooks()
	if err != nil {
		return nil, err
	}

	var resp []response.Get
	for _, book := range books {
		thisBook := bookToResp(book)
		resp = append(resp, thisBook)
	}

	return &resp, nil
}

func bookToResp(book model.Book) response.Get {
	return response.Get{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Genre:  book.Genre,
		Year:   book.Year,
		Tag:    book.Tag,
	}
}
