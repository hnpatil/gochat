package room

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/handlers"
	"github.com/hnpatil/gochat/services"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	svc services.Room
}

func New(svc services.Room) handlers.Room {
	return &handler{svc: svc}
}

func (h *handler) Create(ctx *gofr.Context, req *handlers.CreateRoom) (*entities.Room, error) {
	return h.svc.Create(ctx, &services.CreateRoom{Name: req.Name, UserID: req.UserID, Members: req.Members, RoomID: req.RoomID})
}

func (h *handler) Update(ctx *gofr.Context, req *handlers.UpdateRoom) (*entities.Room, error) {
	return h.svc.Update(ctx, &services.UpdateRoom{RoomID: req.ID, UserID: req.UserID, Name: req.Name})
}

func (h *handler) List(ctx *gofr.Context, req *handlers.ListRooms) ([]*entities.Room, error) {
	return h.svc.List(ctx, &services.ListRoom{UserID: req.UserID, Page: req.Page, Size: req.Size, Include: []string{req.Include}})
}

func (h *handler) Get(ctx *gofr.Context, req *handlers.GetRoom) (*entities.Room, error) {
	return h.svc.Get(ctx, &services.GetRoom{UserID: req.UserID, RoomID: req.ID})
}

func (h *handler) Delete(ctx *gofr.Context, req *handlers.DeleteRoom) error {
	return h.svc.Delete(ctx, &services.DeleteRoom{RoomID: req.ID, UserID: req.UserID})
}
