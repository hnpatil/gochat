package user

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/handlers"
	"github.com/hnpatil/gochat/services"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	svc services.User
}

func New(svc services.User) handlers.User {
	return &handler{
		svc: svc,
	}
}

// @Summary      Create a new user
// @Description  Creates a new user and returns the created user details.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        X-User-ID header string true  "External identifier of the user"
// @Param        user body handlers.UserBody true "User request payload"
// @Success      201 {object} handlers.UserResponse "User successfully created"
// @Failure      400 {object} handlers.ErrorResponse "Invalid request payload"
// @Failure      401 {object} handlers.ErrorResponse "Unauthorized"
// @Failure      500 {object} handlers.ErrorResponse "Internal server error"
// @Router       /v1/users [post]
func (h *handler) Create(ctx *gofr.Context, req *handlers.CreateUser) (*entities.User, error) {
	return h.svc.Create(ctx, &services.CreateUser{ID: req.UserID, Metadata: req.Metadata})
}

// @Summary      Update a user
// @Description  Updates an existing user and returns the updated user details.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        X-User-ID header string true  "External identifier of the user"
// @Param        user body handlers.UserBody true "User request payload"
// @Success      200 {object} handlers.UserResponse "User successfully updated"
// @Failure      400 {object} handlers.ErrorResponse "Invalid request payload"
// @Failure      401 {object} handlers.ErrorResponse "Unauthorized"
// @Failure      404 {object} handlers.ErrorResponse "User not found"
// @Failure      500 {object} handlers.ErrorResponse "Internal server error"
// @Router       /v1/users [patch]
func (h *handler) Update(ctx *gofr.Context, req *handlers.UpdateUser) (*entities.User, error) {
	return h.svc.Update(ctx, &services.UpdateUser{ID: req.UserID, Metadata: req.Metadata})
}

// @Summary      List users
// @Description  Retrieves a paginated list of users.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        X-User-ID header string false "External identifier of the user"
// @Param        page query int false "Page number (default: 1)" example(1)
// @Param        size query int false "Number of users per page (default: 20)" example(20)
// @Success      200 {object} handlers.UsersResponse "Successful response with user list"
// @Failure      400 {object} handlers.ErrorResponse "Invalid request parameters"
// @Failure      401 {object} handlers.ErrorResponse "Unauthorized"
// @Failure      500 {object} handlers.ErrorResponse "Internal server error"
// @Router       /v1/users [get]
func (h *handler) List(ctx *gofr.Context, req *handlers.ListUsers) ([]*entities.User, error) {
	return h.svc.List(ctx, &services.ListUsers{UserID: req.UserID, Page: req.Page, Size: req.Size})
}

// @Summary      Delete a user
// @Description  Deletes a user by their external identifier.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        X-User-ID header string true "External identifier of the user"
// @Success      204 "User successfully deleted"
// @Failure      400 {object} handlers.ErrorResponse "Invalid request parameters"
// @Failure      401 {object} handlers.ErrorResponse "Unauthorized"
// @Failure      404 {object} handlers.ErrorResponse "User not found"
// @Failure      500 {object} handlers.ErrorResponse "Internal server error"
// @Router       /v1/users [delete]
func (h *handler) Delete(ctx *gofr.Context, req *handlers.DeleteUser) error {
	return h.svc.Delete(ctx, &services.DeleteUser{ID: req.UserID})
}
