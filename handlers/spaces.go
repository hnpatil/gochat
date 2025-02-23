package handlers

import (
	"github.com/hnpatil/gochat/entities"
)

// ListSpaces represents a request to list spaces of the user
type ListSpaces struct {
	UserRequest
	UpdatedBefore string `query:"updatedBefore" example:"2025-02-08T14:13:39.080551Z"` // Return spaces updated before this timestamp
}

// SpacesResponse represents the response containing a list of spaces
type SpacesResponse struct {
	Data []*entities.UserSpace `json:"data"` // List of messages
}
