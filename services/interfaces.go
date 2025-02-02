package services

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/entities/roommember"
	"gofr.dev/pkg/gofr"
)

type User interface {
	Create(ctx *gofr.Context, req *CreateUser) (*entities.User, error)
	Update(ctx *gofr.Context, req *UpdateUser) (*entities.User, error)
	Get(ctx *gofr.Context, req *GetUser) (*entities.User, error)
	List(ctx *gofr.Context, req *ListUsers) ([]*entities.User, error)
	Delete(ctx *gofr.Context, req *DeleteUser) error
}

type Room interface {
	Create(ctx *gofr.Context, req *CreateRoom) (*entities.Room, error)
	Update(ctx *gofr.Context, req *UpdateRoom) (*entities.Room, error)
	Get(ctx *gofr.Context, req *GetRoom) (*entities.Room, error)
	List(ctx *gofr.Context, req *ListRoom) ([]*entities.Room, error)
	Delete(ctx *gofr.Context, req *DeleteRoom) error
	ValidateRole(ctx *gofr.Context, roomID string, userID string, roles ...roommember.Role) error
}

type Message interface {
	Create(ctx *gofr.Context, req *CreateMessage) (*entities.Message, error)
	Update(ctx *gofr.Context, req *UpdateMessage) (*entities.Message, error)
	Get(ctx *gofr.Context, req *GetMessage) (*entities.Message, error)
	List(ctx *gofr.Context, req *ListMessage) ([]*entities.Message, error)
	Delete(ctx *gofr.Context, req *DeleteMessage) error
}
