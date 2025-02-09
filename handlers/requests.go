package handlers

import (
	"github.com/hnpatil/gochat/entities/message"
)

type CreateMessage struct {
	UserRequest
	RoomID  string         `path:"roomID" validate:"required"`
	Status  message.Status `json:"status" default:"DRAFT" validate:"oneof=DRAFT SENT"`
	Content string         `json:"content"`
}

type UpdateMessage struct {
	UserRequest
	RoomID    string         `path:"roomID" validate:"required"`
	MessageID string         `path:"messageID" validate:"required,uuid"`
	Status    message.Status `json:"status" default:"DRAFT" validate:"oneof=DRAFT SENT"`
	Content   string         `json:"content"`
}

type GetMessage struct {
	UserRequest
	RoomID    string `path:"roomID" validate:"required,uuid"`
	MessageID string `path:"messageID" validate:"required,uuid"`
}

type DeleteMessage struct {
	UserRequest
	RoomID    string `path:"roomID" validate:"required"`
	MessageID string `path:"messageID" validate:"required,uuid"`
}

type ListMessages struct {
	UserRequest
	RoomID         string `path:"roomID" validate:"required"`
	ModifiedBefore string `query:"modifiedBefore" default:"0001-01-01 00:00:00" validate:"datetime=2006-01-02 15:04:05"`
}

type CreateRoom struct {
	UserRequest
	RoomID  string   `json:"roomID"`
	Name    string   `json:"name"`
	Members []string `json:"members" validate:"required,gt=0"`
}

type UpdateRoom struct {
	UserRequest
	ID   string `path:"id" validate:"required"`
	Name string `json:"name"`
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
