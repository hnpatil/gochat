package handlers

import (
	"github.com/hnpatil/gochat/entities"
)

// MessageBody represents the request body for creating a message
type MessageBody struct {
	Content string `json:"content" example:"Hello"` // Message content
}

// CreateMessage represents a request to create a new message in a room
type CreateMessage struct {
	UserRequest
	MessageBody
	RoomID string `path:"roomID" validate:"required" example:"89e47f30"` // Unique identifier of the room where the message is added
}

// ListMessages represents a request to list messages in a room
type ListMessages struct {
	UserRequest
	RoomID        string `path:"roomID" validate:"required" example:"89e47f30"`        // Unique identifier of the room
	CreatedBefore string `query:"createdBefore" example:"2025-02-08T14:13:39.080551Z"` // Return messages created before this timestamp
}

// MessageResponse represents the response containing a single message
type MessageResponse struct {
	Data *entities.Message `json:"data"` // Message data
}

// MessagesResponse represents the response containing a list of messages
type MessagesResponse struct {
	Data []*entities.Message `json:"data"` // List of messages
}
