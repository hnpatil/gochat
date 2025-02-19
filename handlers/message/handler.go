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

// @Summary      Create a message
// @Description  Creates a message in the specified room and returns the created message.
//
//	The requesting user must be a member of the room and is assigned as the message sender.
//
// @Tags         Messages
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        X-User-ID header string true "External identifier of the user"
// @Param        roomID path string true "Unique identifier of the room where the message is created"
// @Param        message body handlers.MessageBody true "Message creation request payload"
// @Success      201 {object} handlers.MessageResponse "Message successfully created"
// @Failure      400 {object} handlers.ErrorResponse "Invalid request payload"
// @Failure      401 {object} handlers.ErrorResponse "Unauthorized"
// @Failure      403 {object} handlers.ErrorResponse "Forbidden – User is not a member of the room"
// @Failure      404 {object} handlers.ErrorResponse "Room not found"
// @Failure      500 {object} handlers.ErrorResponse "Internal server error"
// @Router       /v1/rooms/{roomID}/messages [post]
func (h *handler) Create(ctx *gofr.Context, req *handlers.CreateMessage) (*entities.Message, error) {
	return h.svc.Create(ctx, &services.CreateMessage{
		UserID:  req.UserID,
		RoomID:  req.RoomID,
		Content: req.Content,
	})
}

// @Summary      List messages
// @Description  Retrieves all SENT messages from a specified room, ordered by creation time in descending order.
//
//	The requesting user must be a member of the room.
//
// @Tags         Messages
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        X-User-ID header string true "External identifier of the requesting user"
// @Param        roomID path string true "Unique identifier of the room from which messages are retrieved"
// @Param        createdBefore query string false "Retrieve messages created before this timestamp (RFC 3339 format)"
// @Success      200 {object} handlers.MessagesResponse "Messages retrieved successfully"
// @Failure      400 {object} handlers.ErrorResponse "Invalid request parameters"
// @Failure      401 {object} handlers.ErrorResponse "Unauthorized"
// @Failure      403 {object} handlers.ErrorResponse "Forbidden – User is not a member of the room"
// @Failure      404 {object} handlers.ErrorResponse "Room not found"
// @Failure      500 {object} handlers.ErrorResponse "Internal server error"
// @Router       /v1/rooms/{roomID}/messages [get]
func (h *handler) List(ctx *gofr.Context, req *handlers.ListMessages) ([]*entities.Message, error) {
	var (
		err           error
		createdBefore time.Time
	)

	if req.CreatedBefore != "" {
		createdBefore, err = time.Parse(time.RFC3339, req.CreatedBefore)
		if err != nil {
			return nil, err
		}
	}

	return h.svc.List(ctx, &services.ListMessage{UserID: req.UserID, RoomID: req.RoomID, CreatedBefore: createdBefore})
}
