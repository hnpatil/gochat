package services

import (
	"github.com/hnpatil/gochat/entities"
	"gofr.dev/pkg/gofr"
)

type Message interface {
	Create(ctx *gofr.Context, req *CreateMessage) (*entities.Message, error)
	List(ctx *gofr.Context, req *ListMessages) ([]*entities.Message, error)
}

type UserSpace interface {
	List(ctx *gofr.Context, req *ListSpaces) ([]*entities.UserSpace, error)
}
