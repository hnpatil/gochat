package space

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/handlers"
	"github.com/hnpatil/gochat/services"
	"gofr.dev/pkg/gofr"
	"time"
)

type handler struct {
	svc services.UserSpace
}

func New(svc services.UserSpace) handlers.UserSpace {
	return &handler{svc: svc}
}

// @Summary      List spaces
// @Description  Retrieves a list of spaces that the requesting user is a member of.
// @Tags         Spaces
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        X-User-ID header string true "External identifier of the user"
// @Param        updatedBefore query string false "Retrieve spaces updated before this timestamp (RFC 3339 format)"
// @Success      200 {object} handlers.SpacesResponse "Successful response with space list"
// @Failure      400 {object} handlers.ErrorResponse "Invalid request parameters"
// @Failure      401 {object} handlers.ErrorResponse "Unauthorized"
// @Failure      500 {object} handlers.ErrorResponse "Internal server error"
// @Router       /v1/spaces [get]
func (h *handler) List(ctx *gofr.Context, req *handlers.ListSpaces) ([]*entities.UserSpace, error) {
	var (
		err           error
		updatedBefore time.Time
	)

	if req.UpdatedBefore != "" {
		updatedBefore, err = time.Parse(time.RFC3339, req.UpdatedBefore)
		if err != nil {
			return nil, err
		}
	}
	return h.svc.List(ctx, &services.ListSpaces{UserID: req.UserID, UpdatedBefore: updatedBefore})
}
