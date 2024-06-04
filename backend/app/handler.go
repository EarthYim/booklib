package app

import (
	"booklib/app/model/req"
)

type booklibService interface {
	Add(req.AddRequest) error
	Edit(req.EditRequest) error
	Delete()
	Get()
	GetById()
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
	var req req.AddRequest
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
