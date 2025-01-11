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
