package room

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/handlers"
	"gofr.dev/pkg/gofr"
)

type handler struct {
}

func New() handlers.Room {
	return &handler{}
}

func (h *handler) Create(ctx *gofr.Context, req *handlers.CreateRoom) (*entities.Room, error) {
	ctx.Debugf("received request: %v", req)
	return nil, handlers.NotImplemented{}
}

func (h *handler) Update(ctx *gofr.Context, req *handlers.UpdateRoom) (*entities.Room, error) {
	ctx.Debugf("received request: %v", req)
	return nil, handlers.NotImplemented{}
}

func (h *handler) List(ctx *gofr.Context, req *handlers.ListRooms) ([]*entities.Room, error) {
	ctx.Debugf("received request: %v", req)
	return nil, handlers.NotImplemented{}
}

func (h *handler) Get(ctx *gofr.Context, req *handlers.GetRoom) (*entities.Room, error) {
	ctx.Debugf("received request: %v", req)
	return nil, handlers.NotImplemented{}
}

func (h *handler) Delete(ctx *gofr.Context, req *handlers.DeleteRoom) error {
	ctx.Debugf("received request: %v", req)
	return handlers.NotImplemented{}
}
