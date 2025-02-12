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
