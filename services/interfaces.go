package services

import (
	"github.com/hnpatil/gochat/entities"
	"gofr.dev/pkg/gofr"
)

type User interface {
	Create(ctx *gofr.Context, req *CreateUser) (*entities.User, error)
	Update(ctx *gofr.Context, req *UpdateUser) (*entities.User, error)
	Get(ctx *gofr.Context, req *GetUser) (*entities.User, error)
	List(ctx *gofr.Context, req *ListUsers) ([]*entities.User, error)
	Delete(ctx *gofr.Context, req *DeleteUser) error
}
