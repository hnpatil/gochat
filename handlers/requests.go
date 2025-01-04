package handlers

import (
	"github.com/google/uuid"
	"github.com/hnpatil/gochat/entities/message"
)

type CreateUser struct {
	ExternalID string `json:"externalID" validate:"required"`
	Name       string `json:"name"`
}

type UpdateUser struct {
	ID   string `path:"id" validate:"required,uuid"`
	Name string `json:"name,omitempty"`
}

type GetUser struct {
	ID string `path:"id" validate:"required,uuid"`
}

type DeleteUser struct {
	ID string `path:"id" validate:"required,uuid"`
}

type ListUsers struct {
	ID         string `query:"id"`
	ExternalID string `query:"externalID"`
}

type CreateMessage struct {
	RoomID  string `path:"roomID" validate:"required,uuid"`
	UserID  uuid.UUID
	Status  message.Status `json:"status" validate:"oneof=DRAFT SENT"`
	Content string         `json:"content"`
}

type UpdateMessage struct {
	RoomID    string         `path:"roomID" validate:"required,uuid"`
	MessageID string         `path:"messageID" validate:"required,uuid"`
	Status    message.Status `json:"status" validate:"oneof=DRAFT SENT"`
	Content   string         `json:"content"`
}

type GetMessage struct {
	RoomID    string `path:"roomID" validate:"required,uuid"`
	MessageID string `path:"messageID" validate:"required,uuid"`
}

type DeleteMessage struct {
	RoomID    string `path:"roomID" validate:"required,uuid"`
	MessageID string `path:"messageID" validate:"required,uuid"`
}

type ListMessages struct {
	RoomID     string `path:"roomID" validate:"required,uuid"`
	Status     string `query:"status" validate:"oneof=DRAFT SENT"`
	SentBefore string `query:"sentBefore" validate:"datetime=2006-01-02 15:04:05"`
}

type CreateRoom struct {
	UserID  uuid.UUID `json:"-"`
	Members []string  `json:"members" validate:"required,gt=0"`
}

type UpdateRoom struct {
	ID   string `path:"id" validate:"required,uuid"`
	Name string `json:"name"`
}

type GetRoom struct {
	ID string `path:"id" validate:"required,uuid"`
}

type DeleteRoom struct {
	ID string `path:"id" validate:"required,uuid"`
}

type ListRooms struct {
	UserID        uuid.UUID `json:"-"`
	ModifiedAfter string    `query:"modifiedAfter" validate:"datetime=2006-01-02 15:04:05"`
}
