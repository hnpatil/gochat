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

// @Summary Create a user
// @Description Create a new user and return the created user.
// @Tags Users
// @Accept json
// @Produce json
// @Security ApiKey
// @Param user body handlers.UserBody true "User Request"
// @Param X-User-ID header string true "External identifier of the user"
// @Success 201 {object} handlers.UserResponse
// @Router /v1/users [post]
func (h *handler) Create(ctx *gofr.Context, req *handlers.CreateUser) (*entities.User, error) {
	return h.svc.Create(ctx, &services.CreateUser{ID: req.UserID, Name: req.Name})
}

// @Summary Update a user
// @Description Update a user and return the updated user.
// @Tags Users
// @Accept json
// @Produce json
// @Security ApiKey
// @Param user body handlers.UserBody true "User Request"
// @Param X-User-ID header string true "External identifier of the user"
// @Success 200 {object} handlers.UserResponse
// @Router /v1/users [patch]
func (h *handler) Update(ctx *gofr.Context, req *handlers.UpdateUser) (*entities.User, error) {
	return h.svc.Update(ctx, &services.UpdateUser{ID: req.UserID, Name: req.Name})
}

// @Summary List users
// @Description Retreive list of all users
// @Tags Users
// @Accept json
// @Produce json
// @Security ApiKey
// @Param X-User-ID header string false "External identifier of the user"
// @Param page query int false "Page number" example(1)
// @Param size query int false "Users per page" example(20)
// @Success 200 {object} handlers.UsersResponse
// @Router /v1/users [get]
func (h *handler) List(ctx *gofr.Context, req *handlers.ListUsers) ([]*entities.User, error) {
	return h.svc.List(ctx, &services.ListUsers{UserID: req.UserID, Page: req.Page, Size: req.Size})
}

// @Summary Delete a user
// @Description Delete a user.
// @Tags Users
// @Accept json
// @Produce json
// @Security ApiKey
// @Param X-User-ID header string true "External identifier of the user"
// @Success 204
// @Router /v1/users [delete]
func (h *handler) Delete(ctx *gofr.Context, req *handlers.DeleteUser) error {
	return h.svc.Delete(ctx, &services.DeleteUser{ID: req.UserID})
}
