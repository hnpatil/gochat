package handlers

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/entities/message"
)

type CreateMessage struct {
	UserRequest
	MessageBody
	RoomID string `path:"roomID" validate:"required"`
}

type UpdateMessage struct {
	UserRequest
	MessageBody
	RoomID    string `path:"roomID" validate:"required"`
	MessageID string `path:"messageID" validate:"required,uuid"`
}

type MessageBody struct {
	Status  message.Status `json:"status" default:"DRAFT" validate:"oneof=DRAFT SENT" example:"SENT"` //Message status
	Content string         `json:"content" example:"Hello"`                                           //Message content
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

type MessageResponse struct {
	Data *entities.Message `json:"data"`
}

type MessagesResponse struct {
	Data []*entities.Message `json:"data"`
}
