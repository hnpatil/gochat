package entities

import (
	"github.com/google/uuid"
	"github.com/hnpatil/gochat/entities/message"
	"github.com/hnpatil/gochat/entities/roommember"
	"time"
)

type Entity struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	DeletedAt  time.Time `json:"deletedAt"`
}

type User struct {
	Entity
	ExternalID string `json:"externalID"`
	Name       string `json:"name"`
}

type Message struct {
	Entity
	RoomID   uuid.UUID      `json:"roomID"`
	SenderID uuid.UUID      `json:"senderID"`
	SentAt   time.Time      `json:"sentAt"`
	Status   message.Status `json:"status"`
	Content  string         `json:"content"`
}

type Room struct {
	Entity
	Name    string        `json:"name"`
	IsGroup bool          `json:"isGroup"`
	Members []*RoomMember `json:"members,omitempty"`
}

type RoomMember struct {
	Entity
	RoomID uuid.UUID       `json:"roomID"`
	UserID uuid.UUID       `json:"userID"`
	Role   roommember.Role `json:"role"`
}
