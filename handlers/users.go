package handlers

type UserRequest struct {
	UserID string `header:"X-User-Id" validate:"required" example:"89e46f30"` // Unique identifier of the requesting user
}
