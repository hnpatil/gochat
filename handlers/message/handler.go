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

// @Summary Create a message
// @Description Create a message in the specified room and return the created message. Calling user should be a member in the room. Calling user is added as message sender.
// @Tags Messages
// @Accept json
// @Produce json
// @Security ApiKey
// @Param message body handlers.MessageBody true "Message Request"
// @Param X-User-ID header string true "External identifier of the user"
// @Param roomID path string true "Room id of the room in which message is added"
// @Success 201 {object} handlers.MessageResponse
// @Router /v1/rooms/{roomID}/messages [post]
func (h *handler) Create(ctx *gofr.Context, req *handlers.CreateMessage) (*entities.Message, error) {
	return h.svc.Create(ctx, &services.CreateMessage{
		UserID:  req.UserID,
		RoomID:  req.RoomID,
		Content: req.Content,
		Status:  req.Status,
	})
}

// @Summary Update a message
// @Description Update a message and return the update message. Message should be a DRAFT and Calling user should be the sender of the message.
// @Tags Messages
// @Accept json
// @Produce json
// @Security ApiKey
// @Param message body handlers.MessageBody true "Message Request"
// @Param X-User-ID header string true "External identifier of the user"
// @Param roomID path string true "Room id of the room in which message is added"
// @Param messageID path string true "Message id"
// @Success 200 {object} handlers.MessageResponse
// @Router /v1/rooms/{roomID}/messages/messageID [patch]
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

// @Summary List messages
// @Description List messages in a room. Returns all SENT messages in the room along with drafts created by the calling user in descending order of modified time. Calling user should be a member in the room.
// @Tags Messages
// @Accept json
// @Produce json
// @Security ApiKey
// @Param X-User-ID header string true "External identifier of the user"
// @Param roomID path string true "Room id of the room in which message is added"
// @Param modifiedBefore query string false  "Modified date time upto which messages should be returned"
// @Success 200 {object} handlers.MessagesResponse
// @Router /v1/rooms/{roomID}/messages [get]
func (h *handler) List(ctx *gofr.Context, req *handlers.ListMessages) ([]*entities.Message, error) {
	modifiedBefore, err := time.Parse(time.DateTime, req.ModifiedBefore)
	if err != nil {
		return nil, err
	}

	return h.svc.List(ctx, &services.ListMessage{UserID: req.UserID, RoomID: req.RoomID, ModifiedBefore: modifiedBefore})
}

// @Summary Delete a message
// @Description Delete a message. Message should be a DRAFT and Calling user should be the sender of the message.
// @Tags Messages
// @Accept json
// @Produce json
// @Security ApiKey
// @Param X-User-ID header string true "External identifier of the user"
// @Param roomID path string true "Room id of the room in which message is added"
// @Param messageID path string true "Message id"
// @Success 204
// @Router /v1/rooms/{roomID}/messages/messageID [delete]
func (h *handler) Delete(ctx *gofr.Context, req *handlers.DeleteMessage) error {
	messageID, err := uuid.Parse(req.MessageID)
	if err != nil {
		return err
	}

	return h.svc.Delete(ctx, &services.DeleteMessage{RoomID: req.RoomID, UserID: req.UserID, MessageID: messageID})
}
