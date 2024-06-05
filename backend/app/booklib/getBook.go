package booklib

import (
	"booklib/app/model/response"
)

func (s *booklibService) Get() []response.Get {
	return []response.Get{}
}

func (s *booklibService) GetByID(id int) response.Get {
	return response.Get{}
}
