package message

import (
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
	})
}

// @Summary List messages
// @Description List messages in a room. Returns all SENT messages in the room along with drafts created by the calling user in descending order of modified time. Calling user should be a member in the room.
// @Tags Messages
// @Accept json
// @Produce json
// @Security ApiKey
// @Param X-User-ID header string true "External identifier of the user"
// @Param roomID path string true "Room id of the room in which message is added"
// @Param createdBefore query string false  "Modified date time upto which messages should be returned"
// @Success 200 {object} handlers.MessagesResponse
// @Router /v1/rooms/{roomID}/messages [get]
func (h *handler) List(ctx *gofr.Context, req *handlers.ListMessages) ([]*entities.Message, error) {
	createdBefore, err := time.Parse(time.RFC3339, req.CreatedBefore)
	if err != nil {
		return nil, err
	}

	return h.svc.List(ctx, &services.ListMessage{UserID: req.UserID, RoomID: req.RoomID, CreatedBefore: createdBefore})
}
