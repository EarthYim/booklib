package booklib

import (
	"booklib/app/model/response"
)

func (s *booklibService) GetBook() []response.Get {
	return []response.Get{}
}

func (s *booklibService) GetBookByID(id int) response.Get {
	return response.Get{}
}
