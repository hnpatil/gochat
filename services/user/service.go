package user

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/pkg/metadata"
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
	usr := &entities.User{
		ID:       req.ID,
		Metadata: req.Metadata,
	}

	return s.repo.Create(ctx, usr)
}

func (s *svc) Update(ctx *gofr.Context, req *services.UpdateUser) (*entities.User, error) {
	usr, err := s.repo.Get(ctx, &entities.User{ID: req.ID})
	if err != nil {
		return nil, err
	}

	usrUpdate := &entities.User{}
	usrUpdate.Metadata, err = metadata.ApplyPatch(usr.Metadata, req.Metadata)
	if err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, &entities.User{ID: req.ID}, usrUpdate)
}

func (s *svc) Get(ctx *gofr.Context, req *services.GetUser) (*entities.User, error) {
	return s.repo.Get(ctx, &entities.User{ID: req.ID})
}

func (s *svc) List(ctx *gofr.Context, req *services.ListUsers) ([]*entities.User, error) {
	repoReq := &repos.UserFilter{Page: req.Page, Size: req.Size}
	if req.UserID != "" {
		repoReq.UserID = []string{req.UserID}
	}
	return s.repo.List(ctx, repoReq)
}

func (s *svc) Delete(ctx *gofr.Context, req *services.DeleteUser) error {
	return s.repo.Delete(ctx, &entities.User{ID: req.ID})
}
