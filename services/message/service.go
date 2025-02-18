package message

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/entities/roommember"
	"github.com/hnpatil/gochat/pkg/id"
	"github.com/hnpatil/gochat/repos"
	"github.com/hnpatil/gochat/services"
	"gofr.dev/pkg/gofr"
)

type svc struct {
	repo repos.Message
	room services.Room
}

func New(repo repos.Message, room services.Room) services.Message {
	return &svc{repo: repo, room: room}
}

func (s *svc) Create(ctx *gofr.Context, req *services.CreateMessage) (*entities.Message, error) {
	_, err := s.room.ValidateRole(ctx, req.RoomID, req.UserID, roommember.RoleAdmin, roommember.RoleMember)
	if err != nil {
		return nil, err
	}

	msg := &entities.Message{
		ID:       id.New(),
		SenderID: &req.UserID,
		RoomID:   req.RoomID,
		Content:  req.Content,
	}

	return s.repo.Create(ctx, msg)
}

func (s *svc) List(ctx *gofr.Context, req *services.ListMessage) ([]*entities.Message, error) {
	_, err := s.room.ValidateRole(ctx, req.RoomID, req.UserID, roommember.RoleAdmin, roommember.RoleMember)
	if err != nil {
		return nil, err
	}

	messages, err := s.repo.List(ctx, &repos.MessageFilter{RoomID: req.RoomID, CreatedBefore: req.CreatedBefore})
	if err != nil {
		return nil, err
	}

	return messages, nil
}
