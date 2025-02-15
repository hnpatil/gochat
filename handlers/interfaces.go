package handlers

import (
	"github.com/hnpatil/gochat/entities"
	"gofr.dev/pkg/gofr"
)

type User interface {
	Create(ctx *gofr.Context, req *CreateUser) (*entities.User, error)
	Update(ctx *gofr.Context, req *UpdateUser) (*entities.User, error)
	List(ctx *gofr.Context, req *ListUsers) ([]*entities.User, error)
	Delete(ctx *gofr.Context, req *DeleteUser) error
}

type Message interface {
	Create(ctx *gofr.Context, req *CreateMessage) (*entities.Message, error)
	List(ctx *gofr.Context, req *ListMessages) ([]*entities.Message, error)
}

type Room interface {
	Create(ctx *gofr.Context, req *CreateRoom) (*entities.Room, error)
	Update(ctx *gofr.Context, req *UpdateRoom) (*entities.Room, error)
	List(ctx *gofr.Context, req *ListRooms) ([]*entities.Room, error)
	Get(ctx *gofr.Context, req *GetRoom) (*entities.Room, error)
}
