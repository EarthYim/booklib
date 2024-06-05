package booklib

import (
	"booklib/app/model"
	"booklib/app/model/request"
)

func (s *booklibService) Edit(req request.Edit) error {
	book := model.Book{
		ID:    req.ID,
		Title: req.Title,
		Autor: req.Autor,
		Genre: req.Genre,
		Year:  req.Year,
		Tag:   req.Tag,
	}
	err := s.EditBook(book)
	if err != nil {
		return err
	}
	return nil
}
