package handlers

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/pkg/metadata"
)

type UserRequest struct {
	UserID string `header:"X-User-Id" validate:"required" example:"89e46f30"` // Unique identifier of the requesting user
}

type UserBody struct {
	Metadata metadata.Metadata `json:"metadata,omitempty"` // Metadata associated with the user
}

type CreateUser struct {
	UserRequest
	UserBody
}

type UpdateUser struct {
	UserRequest
	UserBody
}

type ListUsers struct {
	UserID string `header:"X-User-Id" example:"89e46f30"`  // Unique identifier of the requesting user (optional)
	Page   int    `query:"page" example:"1"`               // Page number for pagination
	Size   int    `query:"size" default:"20" example:"20"` // Number of users per page (default: 20)
}

type DeleteUser struct {
	UserRequest
}

// UserResponse represents the response for a single user
type UserResponse struct {
	Data *entities.User `json:"data"` // User data
}

// UsersResponse represents the response for a list of users
type UsersResponse struct {
	Data []*entities.User `json:"data"` // List of users
}
