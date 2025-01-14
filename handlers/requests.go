package handlers

import (
	"github.com/google/uuid"
	"github.com/hnpatil/gochat/entities/message"
)

type UserRequest struct {
	UserID string `header:"X-User-Id" validate:"required"`
}

type CreateUser struct {
	UserRequest
	Name string `json:"name"`
}

type UpdateUser struct {
	UserRequest
	Name string `json:"name,omitempty"`
}

type ListUsers struct {
	UserID string `header:"X-User-Id"`
	Page   int    `query:"page"`
	Size   int    `query:"size"`
}

type DeleteUser struct {
	UserRequest
}

type CreateMessage struct {
	UserRequest
	RoomID  string         `path:"roomID" validate:"required,uuid"`
	Status  message.Status `json:"status" validate:"oneof=DRAFT SENT"`
	Content string         `json:"content"`
}

type UpdateMessage struct {
	UserRequest
	RoomID    string         `path:"roomID" validate:"required,uuid"`
	MessageID string         `path:"messageID" validate:"required,uuid"`
	Status    message.Status `json:"status" validate:"oneof=DRAFT SENT"`
	Content   string         `json:"content"`
}

type GetMessage struct {
	UserRequest
	RoomID    string `path:"roomID" validate:"required,uuid"`
	MessageID string `path:"messageID" validate:"required,uuid"`
}

type DeleteMessage struct {
	UserRequest
	RoomID    string `path:"roomID" validate:"required,uuid"`
	MessageID string `path:"messageID" validate:"required,uuid"`
}

type ListMessages struct {
	UserRequest
	RoomID     string `path:"roomID" validate:"required,uuid"`
	Status     string `query:"status" validate:"oneof=DRAFT SENT"`
	SentBefore string `query:"sentBefore" validate:"datetime=2006-01-02 15:04:05"`
}

type CreateRoom struct {
	UserRequest
	UserID  uuid.UUID `json:"-"`
	Members []string  `json:"members" validate:"required,gt=0"`
}

type UpdateRoom struct {
	UserRequest
	ID   string `path:"id" validate:"required,uuid"`
	Name string `json:"name"`
}

type GetRoom struct {
	UserRequest
	ID string `path:"id" validate:"required,uuid"`
}

type DeleteRoom struct {
	UserRequest
	ID string `path:"id" validate:"required,uuid"`
}

type ListRooms struct {
	UserRequest
	ModifiedAfter string `query:"modifiedAfter" validate:"datetime=2006-01-02 15:04:05"`
}
