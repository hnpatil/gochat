package handlers

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/pkg/metadata"
)

// CreateRoomBody contains details for creating a room
type CreateRoomBody struct {
	ID       string            `json:"id,omitempty" example:"89e47f30"`                              // Optional unique identifier of the room. A default UID is created if not provided.
	Metadata metadata.Metadata `json:"metadata,omitempty"`                                           // Room metadata
	Members  []string          `json:"members" validate:"required,gt=0" example:"89e46f31,89e46f32"` // List of user IDs to be added as room members.
}

// CreateRoom represents a request to create a new room
type CreateRoom struct {
	UserRequest
	CreateRoomBody
}

// UpdateRoomBody contains details for updating a room
type UpdateRoomBody struct {
	Metadata metadata.Metadata `json:"metadata,omitempty"` // Room metadata
}

// UpdateRoom represents a request to update an existing room
type UpdateRoom struct {
	UserRequest
	UpdateRoomBody
	ID string `path:"id" validate:"required" example:"89e47f30"` // Unique identifier of the room to update
}

// GetRoom represents a request to retrieve a specific room
type GetRoom struct {
	UserRequest
	ID string `path:"id" validate:"required" example:"89e47f30"` // Unique identifier of the room to retrieve
}

// ListRooms represents a request to list rooms the user is a member of
type ListRooms struct {
	UserRequest
	Page    int    `query:"page" example:"1"`               // Page number for pagination
	Size    int    `query:"size" default:"20" example:"20"` // Number of rooms per page (default: 20)
	Include string `query:"include" example:"members"`      // Additional objects to include in the response
}

// RoomResponse represents the response containing a single room
type RoomResponse struct {
	Data *entities.Room `json:"data"` // Room data
}

// RoomsResponse represents the response containing a list of rooms
type RoomsResponse struct {
	Data []*entities.Room `json:"data"` // List of rooms
}
