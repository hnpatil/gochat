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

// @Summary Create a room
// @Description Create a room and return the created room. Calling user is added as an ADMIN in the room.
// @Tags Rooms
// @Accept json
// @Produce json
// @Security ApiKey
// @Param user body handlers.CreateRoomBody true "Room Request"
// @Param X-User-ID header string true "External identifier of the user"
// @Success 201 {object} handlers.RoomResponse
// @Router /v1/rooms [post]
func (h *handler) Create(ctx *gofr.Context, req *handlers.CreateRoom) (*entities.Room, error) {
	return h.svc.Create(ctx, &services.CreateRoom{Name: req.Name, UserID: req.UserID, Members: req.Members, RoomID: req.RoomID})
}

// @Summary Update a room
// @Description Update a room and return the updated room. Calling user should be an ADMIN in the room.
// @Tags Rooms
// @Accept json
// @Produce json
// @Security ApiKey
// @Param user body handlers.UpdateRoomBody true "Room Request"
// @Param X-User-ID header string true "External identifier of the user"
// @Param id path string true "Room ID"
// @Success 200 {object} handlers.RoomResponse
// @Router /v1/rooms/{id} [patch]
func (h *handler) Update(ctx *gofr.Context, req *handlers.UpdateRoom) (*entities.Room, error) {
	return h.svc.Update(ctx, &services.UpdateRoom{RoomID: req.ID, UserID: req.UserID, Name: req.Name})
}

// @Summary List rooms
// @Description Retreive list of all rooms that calling user is a member of.
// @Tags Rooms
// @Accept json
// @Produce json
// @Security ApiKey
// @Param X-User-ID header string true "External identifier of the user"
// @Param id path string true "Room ID"
// @Param page query string false "Page number" example(1)
// @Param size query string false "Rooms per page" example(20)
// @Param include query string false "Additional objects to be included in the response" example(members)
// @Success 200 {object} handlers.RoomsResponse
// @Router /v1/rooms [get]
func (h *handler) List(ctx *gofr.Context, req *handlers.ListRooms) ([]*entities.Room, error) {
	return h.svc.List(ctx, &services.ListRoom{UserID: req.UserID, Page: req.Page, Size: req.Size, Include: []string{req.Include}})
}

// @Summary Get room
// @Description Get a single room by its id.
// @Tags Rooms
// @Accept json
// @Produce json
// @Security ApiKey
// @Param X-User-ID header string true "External identifier of the user"
// @Param id path string true "Room ID"
// @Success 200 {object} handlers.RoomResponse
// @Router /v1/rooms/{id} [get]
func (h *handler) Get(ctx *gofr.Context, req *handlers.GetRoom) (*entities.Room, error) {
	return h.svc.Get(ctx, &services.GetRoom{UserID: req.UserID, RoomID: req.ID})
}

// @Summary Delete a room
// @Description Delete a room. Calling user should be an ADMIN in the room.
// @Tags Rooms
// @Accept json
// @Produce json
// @Security ApiKey
// @Param X-User-ID header string true "External identifier of the user"
// @Param id path string true "Room ID"
// @Success 204
// @Router /v1/rooms/{id} [delete]
func (h *handler) Delete(ctx *gofr.Context, req *handlers.DeleteRoom) error {
	return h.svc.Delete(ctx, &services.DeleteRoom{RoomID: req.ID, UserID: req.UserID})
}
