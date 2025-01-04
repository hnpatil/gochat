package user

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/handlers"
	"gofr.dev/pkg/gofr"
)

type handler struct {
}

func New() handlers.User {
	return &handler{}
}

func (h *handler) Create(ctx *gofr.Context, req *handlers.CreateUser) (*entities.User, error) {
	ctx.Debugf("received request: %v", req)
	return nil, handlers.NotImplemented{}
}

func (h *handler) Update(ctx *gofr.Context, req *handlers.UpdateUser) (*entities.User, error) {
	ctx.Debugf("received request: %v", req)
	return nil, handlers.NotImplemented{}
}

func (h *handler) Get(ctx *gofr.Context, req *handlers.GetUser) (*entities.User, error) {
	ctx.Debugf("received request: %v", req)
	return nil, handlers.NotImplemented{}
}

func (h *handler) List(ctx *gofr.Context, req *handlers.ListUsers) ([]*entities.User, error) {
	ctx.Debugf("received request: %v", req)
	return nil, handlers.NotImplemented{}
}

func (h *handler) Delete(ctx *gofr.Context, req *handlers.DeleteUser) error {
	ctx.Debugf("received request: %v", req)
	return handlers.NotImplemented{}
}
