package services

import (
	"github.com/hnpatil/gochat/pkg/metadata"
	"time"
)

type CreateUser struct {
	ID       string
	Metadata metadata.Metadata
}

type UpdateUser struct {
	ID       string
	Metadata metadata.Metadata
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
	RoomID   string
	UserID   string
	Metadata metadata.Metadata
	Members  []string
}

type UpdateRoom struct {
	UserID   string
	RoomID   string
	Metadata metadata.Metadata
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

type CreateMessage struct {
	UserID  string
	RoomID  string
	Content string
}

type ListMessage struct {
	UserID        string
	RoomID        string
	CreatedBefore time.Time
}
