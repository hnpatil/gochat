package message

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/handlers"
	"gofr.dev/pkg/gofr"
)

type handler struct {
}

func New() handlers.Message {
	return &handler{}
}

func (h *handler) Create(ctx *gofr.Context, req *handlers.CreateMessage) (*entities.Message, error) {
	ctx.Debugf("received request: %v", req)
	return nil, handlers.NotImplemented{}
}

func (h *handler) Update(ctx *gofr.Context, req *handlers.UpdateMessage) (*entities.Message, error) {
	ctx.Debugf("received request: %v", req)
	return nil, handlers.NotImplemented{}
}

func (h *handler) Get(ctx *gofr.Context, req *handlers.GetMessage) (*entities.Message, error) {
	ctx.Debugf("received request: %v", req)
	return nil, handlers.NotImplemented{}
}

func (h *handler) List(ctx *gofr.Context, req *handlers.ListMessages) ([]*entities.Message, error) {
	ctx.Debugf("received request: %v", req)
	return nil, handlers.NotImplemented{}
}

func (h *handler) Delete(ctx *gofr.Context, req *handlers.DeleteMessage) error {
	ctx.Debugf("received request: %v", req)
	return handlers.NotImplemented{}
}
