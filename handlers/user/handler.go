package user

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/handlers"
	"github.com/hnpatil/gochat/services"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	svc services.User
}

func New(svc services.User) handlers.User {
	return &handler{
		svc: svc,
	}
}

func (h *handler) Create(ctx *gofr.Context, req *handlers.CreateUser) (*entities.User, error) {
	return h.svc.Create(ctx, &services.CreateUser{ID: req.UserID, Name: req.Name})
}

func (h *handler) Update(ctx *gofr.Context, req *handlers.UpdateUser) (*entities.User, error) {
	return h.svc.Update(ctx, &services.UpdateUser{ID: req.UserID, Name: req.Name})
}

func (h *handler) List(ctx *gofr.Context, req *handlers.ListUsers) ([]*entities.User, error) {
	return h.svc.List(ctx, &services.ListUsers{UserID: req.UserID, Page: req.Page, Size: req.Size})
}

func (h *handler) Delete(ctx *gofr.Context, req *handlers.DeleteUser) error {
	return h.svc.Delete(ctx, &services.DeleteUser{ID: req.UserID})
}
