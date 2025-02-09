package handlers

import "github.com/hnpatil/gochat/entities"

type UserRequest struct {
	UserID string `header:"X-User-Id" validate:"required"`
}

type CreateUser struct {
	UserRequest
	UserBody
}

type UserBody struct {
	Name string `json:"name,omitempty" example:"John Doe"` //Name of the user
}

type UpdateUser struct {
	UserRequest
	UserBody
}

type ListUsers struct {
	UserID string `header:"X-User-Id"`
	Page   int    `query:"page"`
	Size   int    `query:"size" default:"20"`
}

type DeleteUser struct {
	UserRequest
}

type UserResponse struct {
	Data *entities.User `json:"data"`
}

type UsersResponse struct {
	Data []*entities.User `json:"data"`
}
