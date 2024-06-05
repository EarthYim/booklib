package booklib

import (
	"booklib/app/model"
	"booklib/app/model/request"
)

func (s *booklibService) Add(req request.Add) error {
	//validate request
	book := model.Book{
		Title: req.Title,
		Autor: req.Autor,
		Genre: req.Genre,
		Year:  req.Year,
		Tag:   req.Tag,
	}
	//save to db
	err := s.AddBook(book)
	if err != nil {
		return err
	}
	return nil
}
