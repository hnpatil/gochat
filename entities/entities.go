package entities

import (
	"github.com/google/uuid"
	"github.com/hnpatil/gochat/entities/message"
	"github.com/hnpatil/gochat/entities/roommember"
	"time"
)

type Entity struct {
	CreatedAt  time.Time  `json:"createdAt,omitempty"`
	ModifiedAt time.Time  `json:"modifiedAt,omitempty"`
	DeletedAt  *time.Time `json:"deletedAt,omitempty"`
}

type User struct {
	Entity
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Message struct {
	Entity
	ID       uuid.UUID      `json:"id,omitempty"`
	RoomID   uuid.UUID      `json:"roomID,omitempty"`
	SenderID string         `json:"senderID,omitempty"`
	SentAt   time.Time      `json:"sentAt,omitempty"`
	Status   message.Status `json:"status,omitempty"`
	Content  string         `json:"content,omitempty"`
}

type Room struct {
	Entity
	ID      uuid.UUID     `json:"id,omitempty"`
	Name    string        `json:"name,omitempty"`
	IsGroup bool          `json:"isGroup,omitempty"`
	Members []*RoomMember `json:"members,omitempty,omitempty"`
}

type RoomMember struct {
	Entity
	RoomID uuid.UUID       `json:"roomID,omitempty"`
	UserID string          `json:"userID,omitempty"`
	Role   roommember.Role `json:"role,omitempty"`
}
