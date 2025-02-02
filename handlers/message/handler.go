package message

import (
	"github.com/google/uuid"
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/handlers"
	"github.com/hnpatil/gochat/services"
	"gofr.dev/pkg/gofr"
	"time"
)

type handler struct {
	svc services.Message
}

func New(svc services.Message) handlers.Message {
	return &handler{svc: svc}
}

func (h *handler) Create(ctx *gofr.Context, req *handlers.CreateMessage) (*entities.Message, error) {
	return h.svc.Create(ctx, &services.CreateMessage{
		UserID:  req.UserID,
		RoomID:  req.RoomID,
		Content: req.Content,
		Status:  req.Status,
	})
}

func (h *handler) Update(ctx *gofr.Context, req *handlers.UpdateMessage) (*entities.Message, error) {
	messageID, err := uuid.Parse(req.MessageID)
	if err != nil {
		return nil, err
	}

	return h.svc.Update(ctx, &services.UpdateMessage{
		UserID:    req.UserID,
		RoomID:    req.RoomID,
		Content:   req.Content,
		Status:    req.Status,
		MessageID: messageID,
	})
}

func (h *handler) Get(ctx *gofr.Context, req *handlers.GetMessage) (*entities.Message, error) {
	return nil, handlers.NotImplemented{}
}

func (h *handler) List(ctx *gofr.Context, req *handlers.ListMessages) ([]*entities.Message, error) {
	modifiedBefore, err := time.Parse(time.DateTime, req.ModifiedBefore)
	if err != nil {
		return nil, err
	}

	return h.svc.List(ctx, &services.ListMessage{UserID: req.UserID, RoomID: req.RoomID, ModifiedBefore: modifiedBefore})
}

func (h *handler) Delete(ctx *gofr.Context, req *handlers.DeleteMessage) error {
	messageID, err := uuid.Parse(req.MessageID)
	if err != nil {
		return err
	}

	return h.svc.Delete(ctx, &services.DeleteMessage{RoomID: req.RoomID, UserID: req.UserID, MessageID: messageID})
}
