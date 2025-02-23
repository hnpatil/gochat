package message

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/entities/space"
	"github.com/hnpatil/gochat/errors"
	"github.com/hnpatil/gochat/pkg/id"
	"github.com/hnpatil/gochat/pkg/set"
	"github.com/hnpatil/gochat/repos"
	"github.com/hnpatil/gochat/services"
	"gofr.dev/pkg/gofr"
	"slices"
	"time"
)

type svc struct {
	messageRepo repos.Message
	spaceRepo   repos.UserSpace
}

func New(messageRepo repos.Message, spaceRepo repos.UserSpace) services.Message {
	return &svc{messageRepo: messageRepo, spaceRepo: spaceRepo}
}

func (s *svc) Create(ctx *gofr.Context, req *services.CreateMessage) (*entities.Message, error) {
	recipientSet := set.New[string](req.Recipients...)
	recipientSet.Add(req.UserID)

	recipients := recipientSet.Values()

	msg, err := s.messageRepo.Create(ctx, &entities.Message{
		SpaceID:   id.New(recipients...),
		CreatedAt: time.Now(),
		Data: &entities.MessageData{
			Content:    req.Content,
			SenderID:   req.UserID,
			Recipients: recipients,
		},
	})

	if err != nil {
		return nil, err
	}

	userSpaces := make([]*entities.UserSpace, len(recipients))
	for i, userID := range recipients {
		userSpaces[i] = &entities.UserSpace{
			UserID:    userID,
			SpaceID:   msg.SpaceID,
			UpdatedAt: msg.CreatedAt,
			Data: &entities.UserSpaceData{
				Members: recipients,
				Preview: msg.Data.Content,
			},
		}
	}

	err = s.spaceRepo.UpsertMany(ctx, userSpaces)
	if err != nil {
		ctx.Logger.Errorf("failed to upsert space %s: %s", msg.SpaceID, err)
	}

	return msg, nil
}

func (s *svc) List(ctx *gofr.Context, req *services.ListMessages) ([]*entities.Message, error) {
	messages, err := s.messageRepo.List(ctx, &repos.MessageFilter{
		CreatedBefore: req.CreatedBefore,
		SpaceID:       req.SpaceID,
	})

	if err != nil {
		return nil, err
	}

	if len(messages) > 0 && !slices.Contains(messages[0].Data.Recipients, req.UserID) {
		return nil, errors.UnAuthorised("list messages", space.Entity)
	}

	return messages, nil
}
