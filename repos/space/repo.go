package space

import (
	"encoding/json"
	"github.com/gocql/gocql"
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/entities/space"
	"github.com/hnpatil/gochat/pkg/id"
	"github.com/hnpatil/gochat/repos"
	"github.com/huandu/go-sqlbuilder"
	"gofr.dev/pkg/gofr"
	"time"
)

type repo struct {
	cass *gocql.Session
}

func New(cass *gocql.Session) repos.UserSpace {
	return &repo{cass: cass}
}

func (r *repo) List(_ *gofr.Context, filter *repos.SpaceFilter) ([]*entities.UserSpace, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb = sb.Select(space.FieldUserID, space.FieldSpaceID, space.FieldUpdatedAt, space.FieldData).
		From(space.View).
		Where(sb.Equal(space.FieldUserID, filter.UserID))

	if !filter.UpdatedBefore.IsZero() {
		sb = sb.Where(sb.LessThan(space.FieldUpdatedAt, filter.UpdatedBefore))
	}

	query, args := sb.OrderBy(space.FieldUpdatedAt).Desc().Limit(entities.DefaultSize).Build()
	iter := r.cass.Query(query, args...).Iter()

	defer iter.Close()

	var (
		userID    string
		spaceID   string
		updatedAt time.Time
		data      string
		response  []*entities.UserSpace
	)

	for iter.Scan(&userID, &spaceID, &updatedAt, &data) {
		spaceData := &entities.UserSpaceData{}
		err := json.Unmarshal([]byte(data), spaceData)
		if err != nil {
			return nil, err
		}

		response = append(response, &entities.UserSpace{
			UserID:    userID,
			SpaceID:   id.ID(spaceID),
			UpdatedAt: updatedAt,
			Data:      spaceData,
		})
	}

	return response, nil
}

func (r *repo) UpsertMany(_ *gofr.Context, request []*entities.UserSpace) error {
	batch := r.cass.NewBatch(gocql.UnloggedBatch)

	for _, req := range request {
		data, err := json.Marshal(req.Data)
		if err != nil {
			return err
		}

		query, args := sqlbuilder.NewInsertBuilder().
			InsertInto(space.Table).
			Cols(space.FieldUserID, space.FieldSpaceID, space.FieldUpdatedAt, space.FieldData).
			Values(req.UserID, req.SpaceID, req.UpdatedAt, string(data)).
			Build()

		batch.Query(query, args...)
	}

	return r.cass.ExecuteBatch(batch)
}
