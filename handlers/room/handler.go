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

// @Summary      Create a room
// @Description  Creates a new room and returns the room details. The requesting user is assigned as an ADMIN in the room.
// @Tags         Rooms
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        X-User-ID header string true "External identifier of the user"
// @Param        user body handlers.CreateRoomBody true "Room creation request payload"
// @Success      201 {object} handlers.RoomResponse "Room successfully created"
// @Failure      400 {object} handlers.ErrorResponse "Invalid request payload"
// @Failure      401 {object} handlers.ErrorResponse "Unauthorized"
// @Failure      500 {object} handlers.ErrorResponse "Internal server error"
// @Router       /v1/rooms [post]
func (h *handler) Create(ctx *gofr.Context, req *handlers.CreateRoom) (*entities.Room, error) {
	return h.svc.Create(ctx, &services.CreateRoom{Metadata: req.Metadata, UserID: req.UserID, Members: req.Members, RoomID: req.ID})
}

// @Summary      Update a room
// @Description  Updates an existing room and returns the updated room details. The requesting user must be an ADMIN in the room.
// @Tags         Rooms
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        X-User-ID header string true "External identifier of the user"
// @Param        id path string true "Room ID"
// @Param        user body handlers.UpdateRoomBody true "Room update request payload"
// @Success      200 {object} handlers.RoomResponse "Room successfully updated"
// @Failure      400 {object} handlers.ErrorResponse "Invalid request payload"
// @Failure      401 {object} handlers.ErrorResponse "Unauthorized"
// @Failure      403 {object} handlers.ErrorResponse "Forbidden – User is not an ADMIN"
// @Failure      404 {object} handlers.ErrorResponse "Room not found"
// @Failure      500 {object} handlers.ErrorResponse "Internal server error"
// @Router       /v1/rooms/{id} [patch]
func (h *handler) Update(ctx *gofr.Context, req *handlers.UpdateRoom) (*entities.Room, error) {
	return h.svc.Update(ctx, &services.UpdateRoom{RoomID: req.ID, UserID: req.UserID, Metadata: req.Metadata})
}

// @Summary      List rooms
// @Description  Retrieves a paginated list of rooms that the requesting user is a member of.
// @Tags         Rooms
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        X-User-ID header string true "External identifier of the user"
// @Param        page query int false "Page number (default: 1)" example(1)
// @Param        size query int false "Number of rooms per page (default: 20)" example(20)
// @Param        include query string false "Additional objects to include in the response (e.g., members)" example("members")
// @Success      200 {object} handlers.RoomsResponse "Successful response with room list"
// @Failure      400 {object} handlers.ErrorResponse "Invalid request parameters"
// @Failure      401 {object} handlers.ErrorResponse "Unauthorized"
// @Failure      500 {object} handlers.ErrorResponse "Internal server error"
// @Router       /v1/rooms [get]
func (h *handler) List(ctx *gofr.Context, req *handlers.ListRooms) ([]*entities.Room, error) {
	return h.svc.List(ctx, &services.ListRoom{UserID: req.UserID, Page: req.Page, Size: req.Size, Include: []string{req.Include}})
}

// @Summary      Get room
// @Description  Retrieves details of a single room by its ID.
// @Tags         Rooms
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        X-User-ID header string true "External identifier of the user"
// @Param        id path string true "Unique identifier of the room"
// @Success      200 {object} handlers.RoomResponse "Room details retrieved successfully"
// @Failure      400 {object} handlers.ErrorResponse "Invalid request parameters"
// @Failure      401 {object} handlers.ErrorResponse "Unauthorized"
// @Failure      403 {object} handlers.ErrorResponse "Forbidden – User is not a member of the room"
// @Failure      404 {object} handlers.ErrorResponse "Room not found"
// @Failure      500 {object} handlers.ErrorResponse "Internal server error"
// @Router       /v1/rooms/{id} [get]
func (h *handler) Get(ctx *gofr.Context, req *handlers.GetRoom) (*entities.Room, error) {
	return h.svc.Get(ctx, &services.GetRoom{UserID: req.UserID, RoomID: req.ID})
}
