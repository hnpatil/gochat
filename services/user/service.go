package user

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/repos"
	"github.com/hnpatil/gochat/services"
	"gofr.dev/pkg/gofr"
)

type svc struct {
	repo repos.User
}

func New(repo repos.User) services.User {
	return &svc{repo: repo}
}

func (s *svc) Create(ctx *gofr.Context, req *services.CreateUser) (*entities.User, error) {
	return s.repo.Create(ctx, &entities.User{Name: req.Name, ID: req.ID})
}

func (s *svc) Update(ctx *gofr.Context, req *services.UpdateUser) (*entities.User, error) {
	return s.repo.Update(ctx, &entities.User{ID: req.ID}, &entities.User{Name: req.Name})
}

func (s *svc) Get(ctx *gofr.Context, req *services.GetUser) (*entities.User, error) {
	return s.repo.Get(ctx, &entities.User{ID: req.ID})
}

func (s *svc) List(ctx *gofr.Context, req *services.ListUsers) ([]*entities.User, error) {
	return s.repo.List(ctx, &repos.UserFilter{UserID: []string{req.UserID}, Page: req.Page, Size: req.Size})
}

func (s *svc) Delete(ctx *gofr.Context, req *services.DeleteUser) error {
	return s.repo.Delete(ctx, &entities.User{ID: req.ID})
}
