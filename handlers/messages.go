package handlers

import (
	"github.com/hnpatil/gochat/entities"
)

type CreateMessage struct {
	UserRequest
	MessageBody
	RoomID string `path:"roomID" validate:"required"`
}

type MessageBody struct {
	Content string `json:"content" example:"Hello"` //Message content
}

type GetMessage struct {
	UserRequest
	RoomID    string `path:"roomID" validate:"required,uuid"`
	MessageID string `path:"messageID" validate:"required,uuid"`
}

type ListMessages struct {
	UserRequest
	RoomID        string `path:"roomID" validate:"required"`
	CreatedBefore string `query:"createdBefore"`
}

type MessageResponse struct {
	Data *entities.Message `json:"data"`
}

type MessagesResponse struct {
	Data []*entities.Message `json:"data"`
}
