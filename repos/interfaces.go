package repos

import (
	"github.com/hnpatil/gochat/entities"
	"gofr.dev/pkg/gofr"
)

type User interface {
	Create(ctx *gofr.Context, request *entities.User) (*entities.User, error)
	Update(ctx *gofr.Context, filter, request *entities.User) (*entities.User, error)
	Get(ctx *gofr.Context, filter *entities.User) (*entities.User, error)
	List(ctx *gofr.Context, filter *UserFilter) ([]*entities.User, error)
	Delete(ctx *gofr.Context, filter *entities.User) error
}

type Room interface {
	Create(ctx *gofr.Context, request *entities.Room) (*entities.Room, error)
	Update(ctx *gofr.Context, filter, request *entities.Room) (*entities.Room, error)
	Get(ctx *gofr.Context, filter *entities.Room) (*entities.Room, error)
	List(ctx *gofr.Context, filter *RoomFilter) ([]*entities.Room, error)
}

type Message interface {
	Create(ctx *gofr.Context, request *entities.Message) (*entities.Message, error)
	Get(ctx *gofr.Context, filter *entities.Message) (*entities.Message, error)
	List(ctx *gofr.Context, filter *MessageFilter) ([]*entities.Message, error)
}
