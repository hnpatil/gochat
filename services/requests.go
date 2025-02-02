package services

import (
	"github.com/google/uuid"
	"github.com/hnpatil/gochat/entities/message"
	"time"
)

type CreateUser struct {
	ID   string
	Name string
}

type UpdateUser struct {
	ID   string
	Name string
}

type GetUser struct {
	ID string
}

type DeleteUser struct {
	ID string
}

type ListUsers struct {
	Page   int
	Size   int
	UserID string
}

type CreateRoom struct {
	RoomID  string
	UserID  string
	Name    string
	Members []string
}

type UpdateRoom struct {
	UserID string
	RoomID string
	Name   string
}

type GetRoom struct {
	UserID string
	RoomID string
}

type ListRoom struct {
	Include []string
	UserID  string
	Page    int
	Size    int
}

type DeleteRoom struct {
	RoomID string
	UserID string
}

type CreateMessage struct {
	UserID  string
	RoomID  string
	Status  message.Status
	Content string
}

type UpdateMessage struct {
	UserID    string
	RoomID    string
	MessageID uuid.UUID
	Status    message.Status
	Content   string
}

type GetMessage struct {
}

type ListMessage struct {
	UserID         string
	RoomID         string
	ModifiedBefore time.Time
}

type DeleteMessage struct {
	UserID    string
	RoomID    string
	MessageID uuid.UUID
}
