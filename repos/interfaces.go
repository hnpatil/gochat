package repos

import (
	"github.com/hnpatil/gochat/entities"
	"gofr.dev/pkg/gofr"
)

type Message interface {
	List(ctx *gofr.Context, filter *MessageFilter) ([]*entities.Message, error)
	Create(ctx *gofr.Context, request *entities.Message) (*entities.Message, error)
}

type UserSpace interface {
	List(ctx *gofr.Context, filter *SpaceFilter) ([]*entities.UserSpace, error)
	UpsertMany(ctx *gofr.Context, request []*entities.UserSpace) error
}
