package app

import (
	"booklib/app/model/request"
	"booklib/app/model/response"
)

type booklibService interface {
	Add(request.Add) error
	Edit(request.Edit) error
	Delete(request.Delete) error
	Get() []response.Get
	GetById() response.Get
}

type bookHandler struct {
	booklibService
}

func NewBookHandler(service booklibService) *bookHandler {
	return &bookHandler{
		booklibService: service,
	}
}

func (h *bookHandler) HandleAdd(ctx Context) {
	var req request.Add
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = h.Add(req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON("success")
}
