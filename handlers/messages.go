package handlers

import (
	"github.com/hnpatil/gochat/entities"
)

// MessageBody represents a send message request body
type MessageBody struct {
	Content    string   `json:"content" example:"Hello"`                                         // Message content
	Recipients []string `json:"recipients" validate:"required,gt=0" example:"89e46f31,89e46f32"` // User IDs to be added as recipients.
}

// CreateMessage represents a request to create a new message in a room
type CreateMessage struct {
	UserRequest
	MessageBody
}

// ListMessages represents a request to list messages in a room
type ListMessages struct {
	UserRequest
	SpaceID       string `header:"X-Space-Id" validate:"required" example:"89e46f30"`  // Unique identifier of the space
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
