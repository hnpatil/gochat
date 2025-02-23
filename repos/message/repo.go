package message

import (
	"encoding/json"
	"github.com/gocql/gocql"
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/entities/message"
	"github.com/hnpatil/gochat/pkg/id"
	"github.com/hnpatil/gochat/repos"
	"github.com/huandu/go-sqlbuilder"
	"gofr.dev/pkg/gofr"
	"time"
)

type repo struct {
	cass *gocql.Session
}

func New(cass *gocql.Session) repos.Message {
	return &repo{cass: cass}
}

func (r *repo) List(_ *gofr.Context, filter *repos.MessageFilter) ([]*entities.Message, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb = sb.
		Select(message.FieldSpaceID, message.FieldCreatedAt, message.FieldData).
		From(message.Table).
		Where(sb.Equal(message.FieldSpaceID, filter.SpaceID))

	if !filter.CreatedBefore.IsZero() {
		sb = sb.Where(sb.LessThan(message.FieldCreatedAt, filter.CreatedBefore))
	}

	query, args := sb.OrderBy(message.FieldCreatedAt).Desc().Limit(entities.DefaultSize).Build()

	iter := r.cass.Query(query, args...).Iter()

	defer iter.Close()

	var (
		spaceID   string
		data      string
		createdAt time.Time
		response  []*entities.Message
	)

	for iter.Scan(&spaceID, &createdAt, &data) {
		messageData := &entities.MessageData{}
		err := json.Unmarshal([]byte(data), messageData)
		if err != nil {
			return nil, err
		}

		response = append(response, &entities.Message{
			SpaceID:   id.ID(spaceID),
			CreatedAt: createdAt,
			Data:      messageData,
		})
	}

	return response, nil
}

func (r *repo) Create(_ *gofr.Context, request *entities.Message) (*entities.Message, error) {
	data, err := json.Marshal(request.Data)
	if err != nil {
		return nil, err
	}

	query, args := sqlbuilder.NewInsertBuilder().
		InsertInto(message.Table).
		Cols(message.FieldSpaceID, message.FieldCreatedAt, message.FieldData).
		Values(request.SpaceID, request.CreatedAt, string(data)).
		Build()

	err = r.cass.Query(query, args...).Exec()
	if err != nil {
		return nil, err
	}

	return request, nil
}
