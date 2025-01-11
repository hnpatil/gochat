package user

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/repos"
	"github.com/hnpatil/gochat/usecases"
	"gofr.dev/pkg/gofr"
)

type svc struct {
	repo repos.User
}

func New(repo repos.User) usecases.User {
	return &svc{repo: repo}
}

func (s *svc) Create(ctx *gofr.Context, req *usecases.CreateUser) (*entities.User, error) {
	return s.repo.Create(ctx, &entities.User{Name: req.Name, ID: req.ID})
}

func (s *svc) Update(ctx *gofr.Context, req *usecases.UpdateUser) (*entities.User, error) {
	return s.repo.Update(ctx, &entities.User{ID: req.ID}, &entities.User{Name: req.Name})
}

func (s *svc) Get(ctx *gofr.Context, req *usecases.GetUser) (*entities.User, error) {
	return s.repo.Get(ctx, &entities.User{ID: req.ID})
}

func (s *svc) List(ctx *gofr.Context, req *usecases.ListUsers) ([]*entities.User, error) {
	return s.repo.List(ctx, &repos.UserFilter{UserID: req.UserID, Page: req.Page, Size: req.Size})
}

func (s *svc) Delete(ctx *gofr.Context, req *usecases.DeleteUser) error {
	return s.repo.Delete(ctx, &entities.User{ID: req.ID})
}
