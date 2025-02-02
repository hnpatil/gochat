package message

import (
	"database/sql"
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/entities/message"
	"github.com/hnpatil/gochat/repos"
	"github.com/huandu/go-sqlbuilder"
	"gofr.dev/pkg/gofr"
	"time"
)

type repo struct {
}

func New() repos.Message {
	return &repo{}
}

func (r *repo) Create(ctx *gofr.Context, request *entities.Message) (*entities.Message, error) {
	query, args := sqlbuilder.NewInsertBuilder().
		InsertInto(message.Table).
		Cols(message.FieldID, message.FieldRoomID, message.FieldSenderID, message.FieldSentAt, message.FieldStatus, message.FieldContent).
		Values(request.ID, request.RoomID, request.SenderID, request.SentAt, request.Status, request.Content).
		Build()

	_, err := ctx.SQL.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, &entities.Message{ID: request.ID})
}

func (r *repo) Update(ctx *gofr.Context, filter, request *entities.Message) (*entities.Message, error) {
	sb := sqlbuilder.NewUpdateBuilder()
	sets := []string{sb.Assign(message.FieldModifiedAt, time.Now())}

	if request.Content != "" {
		sets = append(sets, sb.Assign(message.FieldContent, request.Content))
	}

	if request.Status != "" {
		sets = append(sets, sb.Assign(message.FieldStatus, request.Status))
	}

	if request.SentAt != nil {
		sets = append(sets, sb.Assign(message.FieldSentAt, request.SentAt))
	}

	query, args := sb.Update(message.Table).Set(sets...).Where(sb.Equal(message.FieldID, filter.ID)).Build()
	res, err := ctx.SQL.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if n == 0 {
		return nil, sql.ErrNoRows
	}

	return r.Get(ctx, &entities.Message{ID: filter.ID})
}

func (r *repo) Get(ctx *gofr.Context, filter *entities.Message) (*entities.Message, error) {
	sb := sqlbuilder.NewSelectBuilder()
	query, args := sb.
		Select(
			message.FieldID, message.FieldRoomID, message.FieldSenderID, message.FieldSentAt, message.FieldStatus,
			message.FieldContent, message.FieldCreatedAt, message.FieldModifiedAt, message.FieldDeletedAt,
		).
		From(message.Table).Where(sb.Equal(message.FieldID, filter.ID)).Build()

	row := ctx.SQL.QueryRowContext(ctx, query, args...)
	if err := row.Err(); err != nil {
		return nil, err
	}

	msg := &entities.Message{}

	err := row.Scan(
		&msg.ID, &msg.RoomID, &msg.SenderID, &msg.SentAt, &msg.Status,
		&msg.Content, &msg.CreatedAt, &msg.ModifiedAt, &msg.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (r *repo) List(ctx *gofr.Context, filter *repos.MessageFilter) ([]*entities.Message, error) {
	sb := sqlbuilder.NewSelectBuilder().
		Select(
			message.FieldID, message.FieldRoomID, message.FieldSenderID, message.FieldSentAt, message.FieldStatus,
			message.FieldContent, message.FieldCreatedAt, message.FieldModifiedAt, message.FieldDeletedAt,
		).
		From(message.Table)

	if filter.RoomID != "" {
		sb = sb.Where(sb.EQ(message.FieldRoomID, filter.RoomID))
	}

	if !filter.ModifiedBefore.IsZero() {
		sb = sb.Where(sb.LessThan(message.FieldModifiedAt, filter.ModifiedBefore))
	}

	query, args := sb.OrderBy(message.FieldModifiedAt).Desc().Limit(20).Build()
	rows, err := ctx.SQL.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	msgs := make([]*entities.Message, 0, 20)

	for rows.Next() {
		msg := &entities.Message{}

		err = rows.Scan(
			&msg.ID, &msg.RoomID, &msg.SenderID, &msg.SentAt, &msg.Status,
			&msg.Content, &msg.CreatedAt, &msg.ModifiedAt, &msg.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		msgs = append(msgs, msg)
	}

	return msgs, err
}

func (r *repo) Delete(ctx *gofr.Context, filter *entities.Message) error {
	sb := sqlbuilder.NewDeleteBuilder()
	query, args := sb.DeleteFrom(message.Table).Where(sb.Equal(message.FieldID, filter.ID)).Build()

	res, err := ctx.SQL.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return sql.ErrNoRows
	}

	return nil
}
