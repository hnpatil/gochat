package user

import (
	"github.com/hnpatil/gochat/auth"
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/handlers"
	"github.com/hnpatil/gochat/usecases"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	svc usecases.User
}

func New(svc usecases.User) handlers.User {
	return &handler{
		svc: svc,
	}
}

func (h *handler) Create(ctx *gofr.Context, req *handlers.CreateUser) (*entities.User, error) {
	userID := ctx.Value(auth.UserID).(string)
	return h.svc.Create(ctx, &usecases.CreateUser{ID: userID, Name: req.Name})
}

func (h *handler) Update(ctx *gofr.Context, req *handlers.UpdateUser) (*entities.User, error) {
	userID := ctx.Value(auth.UserID).(string)
	return h.svc.Update(ctx, &usecases.UpdateUser{ID: userID, Name: req.Name})
}

func (h *handler) List(ctx *gofr.Context, req *handlers.ListUsers) ([]*entities.User, error) {
	userID := ctx.Value(auth.UserID).(string)
	return h.svc.List(ctx, &usecases.ListUsers{UserID: userID, Page: req.Page, Size: req.Size})
}

func (h *handler) Delete(ctx *gofr.Context) error {
	userID := ctx.Value(auth.UserID).(string)
	return h.svc.Delete(ctx, &usecases.DeleteUser{ID: userID})
}
