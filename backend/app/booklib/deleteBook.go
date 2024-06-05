package booklib

import (
	"booklib/app/model"
	"booklib/app/model/request"
)

func (s *booklibService) Delete(req request.Delete) error {
	book := model.Book{
		ID: req.ID,
	}
	err := s.DeleteBook(book)
	if err != nil {
		return err
	}
	return nil
}
