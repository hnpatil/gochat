package message

import (
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/entities/message"
	"github.com/hnpatil/gochat/repos"
	"github.com/huandu/go-sqlbuilder"
	"gofr.dev/pkg/gofr"
)

type repo struct {
}

func New() repos.Message {
	return &repo{}
}

func (r *repo) Create(ctx *gofr.Context, request *entities.Message) (*entities.Message, error) {
	query, args := sqlbuilder.NewInsertBuilder().
		InsertInto(message.Table).
		Cols(message.FieldID, message.FieldRoomID, message.FieldSenderID, message.FieldContent).
		Values(request.ID, request.RoomID, request.SenderID, request.Content).
		Build()

	_, err := ctx.SQL.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, repos.Error(err, message.Entity)
	}

	return r.Get(ctx, &entities.Message{RoomID: request.RoomID, ID: request.ID})
}

func (r *repo) Get(ctx *gofr.Context, filter *entities.Message) (*entities.Message, error) {
	sb := sqlbuilder.NewSelectBuilder()
	query, args := sb.
		Select(
			message.FieldID, message.FieldRoomID, message.FieldSenderID,
			message.FieldContent, message.FieldCreatedAt, message.FieldModifiedAt,
		).
		From(message.Table).
		Where(sb.Equal(message.FieldRoomID, filter.RoomID), sb.Equal(message.FieldID, filter.ID)).Build()

	row := ctx.SQL.QueryRowContext(ctx, query, args...)
	if err := row.Err(); err != nil {
		return nil, repos.Error(err, message.Entity)
	}

	msg := &entities.Message{}

	err := row.Scan(&msg.ID, &msg.RoomID, &msg.SenderID, &msg.Content, &msg.CreatedAt, &msg.ModifiedAt)
	if err != nil {
		return nil, repos.Error(err, message.Entity)
	}

	return msg, nil
}

func (r *repo) List(ctx *gofr.Context, filter *repos.MessageFilter) ([]*entities.Message, error) {
	sb := sqlbuilder.NewSelectBuilder().
		Select(
			message.FieldID, message.FieldRoomID, message.FieldSenderID,
			message.FieldContent, message.FieldCreatedAt, message.FieldModifiedAt,
		).
		From(message.Table)

	if filter.RoomID != "" {
		sb = sb.Where(sb.EQ(message.FieldRoomID, filter.RoomID))
	}

	if !filter.CreatedBefore.IsZero() {
		sb = sb.Where(sb.LessThan(message.FieldCreatedAt, filter.CreatedBefore))
	}

	query, args := sb.OrderBy(message.FieldModifiedAt).Desc().Limit(20).Build()
	rows, err := ctx.SQL.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, repos.Error(err, message.Entity)
	}

	defer rows.Close()

	msgs := make([]*entities.Message, 0, 20)

	for rows.Next() {
		msg := &entities.Message{}

		err = rows.Scan(&msg.ID, &msg.RoomID, &msg.SenderID, &msg.Content, &msg.CreatedAt, &msg.ModifiedAt)
		if err != nil {
			return nil, repos.Error(err, message.Entity)
		}

		msgs = append(msgs, msg)
	}

	return msgs, repos.Error(err, message.Entity)
}
