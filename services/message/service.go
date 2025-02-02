package message

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/entities/message"
	"github.com/hnpatil/gochat/entities/roommember"
	"github.com/hnpatil/gochat/repos"
	"github.com/hnpatil/gochat/services"
	"gofr.dev/pkg/gofr"
	"time"
)

type svc struct {
	repo repos.Message
	room services.Room
}

func New(repo repos.Message, room services.Room) services.Message {
	return &svc{repo: repo, room: room}
}

func (s *svc) Create(ctx *gofr.Context, req *services.CreateMessage) (*entities.Message, error) {
	err := s.room.ValidateRole(ctx, req.RoomID, req.UserID, roommember.RoleAdmin, roommember.RoleMember)
	if err != nil {
		return nil, err
	}

	msg := &entities.Message{
		ID:       uuid.New(),
		SenderID: req.UserID,
		RoomID:   req.RoomID,
		Content:  req.Content,
		Status:   req.Status,
	}

	if msg.Status == message.StatusSent {
		now := time.Now()
		msg.SentAt = &now
	}

	return s.repo.Create(ctx, msg)
}

func (s *svc) Update(ctx *gofr.Context, req *services.UpdateMessage) (*entities.Message, error) {
	msg, err := s.repo.Get(ctx, &entities.Message{ID: req.MessageID})
	if err != nil {
		return nil, err
	}

	if msg.SenderID != req.UserID {
		return nil, services.UnAuthorisedError(fmt.Sprintf("user %s cannot edit messsage %s", req.UserID, msg.ID))
	}

	if msg.Status == message.StatusSent {
		return nil, services.ForbiddenError("cannot edit sent message")
	}

	updateReq := &entities.Message{
		Content: req.Content,
		Status:  req.Status,
	}

	if req.Status == message.StatusSent {
		now := time.Now()
		updateReq.SentAt = &now
	}

	return s.repo.Update(ctx, &entities.Message{ID: req.MessageID}, updateReq)
}

func (s *svc) Get(ctx *gofr.Context, req *services.GetMessage) (*entities.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (s *svc) List(ctx *gofr.Context, req *services.ListMessage) ([]*entities.Message, error) {
	err := s.room.ValidateRole(ctx, req.RoomID, req.UserID, roommember.RoleAdmin, roommember.RoleMember)
	if err != nil {
		return nil, err
	}

	messages, err := s.repo.List(ctx, &repos.MessageFilter{RoomID: req.RoomID, ModifiedBefore: req.ModifiedBefore})
	if err != nil {
		return nil, err
	}

	userMessages := make([]*entities.Message, 0, len(messages))

	for _, msg := range messages {
		if msg.Status == message.StatusDraft && msg.SenderID != req.UserID {
			continue
		}

		userMessages = append(userMessages, msg)
	}

	return userMessages, nil
}

func (s *svc) Delete(ctx *gofr.Context, req *services.DeleteMessage) error {
	msg, err := s.repo.Get(ctx, &entities.Message{ID: req.MessageID})
	if err != nil {
		return err
	}

	if msg.SenderID != req.UserID {
		return services.UnAuthorisedError(fmt.Sprintf("user %s cannot delete messsage %s", req.UserID, msg.ID))
	}

	if msg.Status == message.StatusSent {
		return services.ForbiddenError("cannot delete sent message")
	}

	return s.repo.Delete(ctx, &entities.Message{ID: req.MessageID})
}
