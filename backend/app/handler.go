package app

import (
	"booklib/app/model/request"
	"booklib/app/model/response"
)

type booklibService interface {
	Add(request.Add) error
	Edit(request.Edit) error
	Delete(request.Delete) error
	GetAll() (*[]response.Get, error)
	GetByID(id int) (*response.Get, error)
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
	err := ctx.ShouldBindJSON(&req)
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

func (h *bookHandler) HandleEdit(ctx Context) {
	var req request.Edit
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = h.Edit(req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON("success")
}

func (h *bookHandler) HandleDelete(ctx Context) {
	var req request.Delete
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = h.Delete(req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON("success")
}

func (h *bookHandler) HandleGet(ctx Context) {
	resp, err := h.GetAll()
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(resp)
}

func (h *bookHandler) HandleGetByID(ctx Context) {
	var req request.Get
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	resp, err := h.GetByID(req.ID)
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(resp)
}
