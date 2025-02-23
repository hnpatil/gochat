package space

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/repos"
	"github.com/hnpatil/gochat/services"
	"gofr.dev/pkg/gofr"
)

type svc struct {
	repo repos.UserSpace
}

func New(repo repos.UserSpace) services.UserSpace {
	return &svc{repo: repo}
}

func (s *svc) List(ctx *gofr.Context, req *services.ListSpaces) ([]*entities.UserSpace, error) {
	return s.repo.List(ctx, &repos.SpaceFilter{UpdatedBefore: req.UpdatedBefore, UserID: req.UserID})
}
