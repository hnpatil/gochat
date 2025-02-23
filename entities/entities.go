package entities

import (
	"github.com/hnpatil/gochat/pkg/id"
	"time"
)

// Message represents a chat message.
//
// swagger:model
type Message struct {
	// Unique identifier for the space where the message is sent.
	SpaceID id.ID `json:"spaceID" example:"550e8400-e29b-41d4-a716-446655440000"`
	// Timestamp of when the message was created.
	CreatedAt time.Time `json:"createdAt" example:"2025-02-22T14:00:00Z"`
	// Message details including content and sender info.
	Data *MessageData `json:"data"`
}

// MessageData holds the details of a message.
//
// swagger:model
type MessageData struct {
	// The actual content of the message.
	Content string `json:"content" example:"Hello, world!"`
	// The ID of the sender.
	SenderID string `json:"senderID" example:"user_123"`
	// List of recipient user IDs.
	Recipients []string `json:"recipients" example:"user_456,user_789"`
}

// UserSpace represents a user's participation in a space.
//
// swagger:model
type UserSpace struct {
	// Unique identifier for the user.
	UserID string `json:"userID" example:"user_123"`
	// Unique identifier for the space.
	SpaceID id.ID `json:"spaceID" example:"550e8400-e29b-41d4-a716-446655440000"`
	// Timestamp of the last update in the user-space relationship.
	UpdatedAt time.Time `json:"updatedAt" example:"2025-02-22T14:00:00Z"`
	// Additional user-space metadata.
	Data *UserSpaceData `json:"data"`
}

// UserSpaceData contains additional metadata about a user's space participation.
//
// swagger:model
type UserSpaceData struct {
	// A preview message for the space.
	Preview string `json:"preview" example:"Latest message preview here..."`
	// List of user IDs who are members of the space.
	Members []string `json:"members" example:"user_456,user_789"`
}
