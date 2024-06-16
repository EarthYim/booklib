package booklib

import (
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

// func (s *booklibService) GetByName(name string) []response.Get {
// 	books, err := s.GetBookByName(name)
// 	if err != nil {
// 		return []response.Get{}
// 	}
// 	var resp []response.Get
// 	for _, book := range books {
// 		resp = append(resp, response.Get{
// 			ID:    book.ID,
// 			Title: book.Title,
// 			Author: book.Author,
// 			Genre: book.Genre,
// 			Year:  book.Year,
// 			Tag:   book.Tag,
// 		})
// 	}
// 	return resp
// }

// func (s *booklibService) GetByYear(year string) []response.Get {
// 	books, err := s.GetBookByYear(year)
// 	if err != nil {
// 		return []response.Get{}
// 	}
// 	var resp []response.Get
// 	for _, book := range books {
// 		resp = append(resp, response.Get{
// 			ID:    book.ID,
// 			Title: book.Title,
// 			Author: book.Author,
// 			Genre: book.Genre,
// 			Year:  book.Year,
// 			Tag:   book.Tag,
// 		})
// 	}
// 	return resp
// }
