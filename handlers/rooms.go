package handlers

import "github.com/hnpatil/gochat/entities"

type CreateRoom struct {
	UserRequest
	CreateRoomBody
}

type CreateRoomBody struct {
	RoomID  string   `json:"roomID" example:"89e47f30"`                                    //Optional unique identifier of the room. A default UID is created if not present
	Name    string   `json:"name" example:"Friends"`                                       //Optional group name
	Members []string `json:"members" validate:"required,gt=0" example:"89e46f31,89e46f32"` //List of user ids of room members.
}

type UpdateRoom struct {
	UserRequest
	UpdateRoomBody
	ID string `path:"id" validate:"required"`
}

type UpdateRoomBody struct {
	Name string `json:"name" example:"Friends"` //Updated group name
}

type GetRoom struct {
	UserRequest
	ID string `path:"id" validate:"required"`
}

type DeleteRoom struct {
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
