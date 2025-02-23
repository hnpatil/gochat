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
// @Description  Creates a message into space identified by recipients and returns the created message.
//
//	The requesting user is assigned as the message sender.
//
// @Tags         Messages
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        X-User-ID header string true "External identifier of the user"
// @Param        message body handlers.MessageBody true "Message creation request payload"
// @Success      201 {object} handlers.MessageResponse "Message successfully created"
// @Failure      400 {object} handlers.ErrorResponse "Invalid request payload"
// @Failure      401 {object} handlers.ErrorResponse "Unauthorized"
// @Failure      500 {object} handlers.ErrorResponse "Internal server error"
// @Router       /v1/messages [post]
func (h *handler) Create(ctx *gofr.Context, req *handlers.CreateMessage) (*entities.Message, error) {
	return h.svc.Create(ctx, &services.CreateMessage{
		UserID:     req.UserID,
		Content:    req.Content,
		Recipients: req.Recipients,
	})
}

// @Summary      List messages
// @Description  Retrieves all SENT messages from a specified space, ordered by creation time in descending order.
//
//	The requesting user must be a member in the space.
//
// @Tags         Messages
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        X-User-ID header string true "External identifier of the requesting user"
// @Param        X-Space-ID header string true "Unique identifier of the space from which messages are retrieved"
// @Param        createdBefore query string false "Retrieve messages created before this timestamp (RFC 3339 format)"
// @Success      200 {object} handlers.MessagesResponse "Messages retrieved successfully"
// @Failure      400 {object} handlers.ErrorResponse "Invalid request parameters"
// @Failure      401 {object} handlers.ErrorResponse "Unauthorized"
// @Failure      403 {object} handlers.ErrorResponse "Forbidden – User cannot list messages on UserSpace"
// @Failure      404 {object} handlers.ErrorResponse "UserSpace not found"
// @Failure      500 {object} handlers.ErrorResponse "Internal server error"
// @Router       /v1/messages [get]
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

	return h.svc.List(ctx, &services.ListMessages{UserID: req.UserID, SpaceID: req.SpaceID, CreatedBefore: createdBefore})
}
