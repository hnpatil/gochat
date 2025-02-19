package entities

import (
	"github.com/hnpatil/gochat/entities/roommember"
	"github.com/hnpatil/gochat/pkg/id"
	"github.com/hnpatil/gochat/pkg/metadata"
	"time"
)

type Entity struct {
	CreatedAt  time.Time `json:"createdAt,omitempty" example:"2025-02-08T14:13:39.080551Z"`  // Timestamp when the entity was created
	ModifiedAt time.Time `json:"modifiedAt,omitempty" example:"2025-02-08T14:13:39.080551Z"` // Timestamp when the entity was last modified
}

type User struct {
	Entity
	ID       string            `json:"id,omitempty" example:"89e46f30"` // Unique identifier of the user
	Metadata metadata.Metadata `json:"metadata,omitempty"`              // Metadata associated with the user
}

type Message struct {
	Entity
	ID       id.ID   `json:"id,omitempty" example:"89e48f30"`       // Unique identifier of the message
	RoomID   string  `json:"roomID,omitempty" example:"89e47f30"`   // Unique identifier of the room the message belongs to
	SenderID *string `json:"senderID,omitempty" example:"89e46f30"` // Unique identifier of the user who created the message
	Content  string  `json:"content,omitempty" example:"Hello"`     // Message content
}

type Room struct {
	Entity
	ID       string            `json:"id,omitempty" example:"89e47f30"` // Unique identifier of the room
	Members  []*RoomMember     `json:"members,omitempty"`               // List of room members
	Metadata metadata.Metadata `json:"metadata,omitempty"`              // Metadata associated with the room
}

type RoomMember struct {
	Entity
	RoomID string          `json:"roomID,omitempty" example:"89e47f30"` // Unique identifier of the room
	UserID string          `json:"userID,omitempty" example:"89e46f30"` // Unique identifier of the user
	Role   roommember.Role `json:"role,omitempty" example:"ADMIN"`      // Role defining user permissions in the room
	User   *User           `json:"user,omitempty"`                      // User object associated with the room member
}
