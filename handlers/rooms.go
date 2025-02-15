package handlers

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/pkg/metadata"
)

type CreateRoom struct {
	UserRequest
	CreateRoomBody
}

type CreateRoomBody struct {
	ID       string            `json:"id" example:"89e47f30"`                                        //Optional unique identifier of the room. A default UID is created if not present
	Metadata metadata.Metadata `json:"metadata,omitempty"`                                           //Room metadata
	Members  []string          `json:"members" validate:"required,gt=0" example:"89e46f31,89e46f32"` //List of user ids of room members.
}

type UpdateRoom struct {
	UserRequest
	UpdateRoomBody
	ID string `path:"id" validate:"required"`
}

type UpdateRoomBody struct {
	Metadata metadata.Metadata `json:"metadata,omitempty"` //Room metadata
}

type GetRoom struct {
	UserRequest
	ID string `path:"id" validate:"required"`
}

type ListRooms struct {
	UserRequest
	Page    int    `query:"page"`
	Size    int    `query:"size" default:"20"`
	Include string `query:"include"`
}

type RoomResponse struct {
	Data *entities.Room `json:"data"`
}

type RoomsResponse struct {
	Data []*entities.Room `json:"data"`
}
