package room

import (
	"database/sql"
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/entities/room"
	"github.com/hnpatil/gochat/entities/roommember"
	"github.com/hnpatil/gochat/repos"
	"github.com/huandu/go-sqlbuilder"
	"gofr.dev/pkg/gofr"
	"slices"
	"time"
)

type roomRepo struct {
}

func New() repos.Room {
	return &roomRepo{}
}

func (r *roomRepo) Create(ctx *gofr.Context, request *entities.Room) (*entities.Room, error) {
	tx, err := ctx.SQL.Begin()
	if err != nil {
		return nil, err
	}

	query, args := sqlbuilder.NewInsertBuilder().
		InsertInto(room.Table).
		Cols(room.FieldID, room.FieldName, room.FieldIsGroup).
		Values(request.ID, request.Name, request.IsGroup).
		Build()

	_, err = tx.ExecContext(ctx.Context, query, args...)
	if err != nil {
		return nil, err
	}

	for _, member := range request.Members {
		query, args = sqlbuilder.NewInsertBuilder().
			InsertInto(roommember.Table).
			Cols(roommember.FieldRoomID, roommember.FieldUserID, roommember.FieldRole).
			Values(member.RoomID, member.UserID, member.Role).
			Build()

		_, err = tx.ExecContext(ctx.Context, query, args...)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, &entities.Room{ID: request.ID})
}

func (r *roomRepo) Update(ctx *gofr.Context, filter, request *entities.Room) (*entities.Room, error) {
	sb := sqlbuilder.NewUpdateBuilder()
	sets := []string{sb.Assign(room.FieldModifiedAt, time.Now())}

	if request.Name != "" {
		sets = append(sets, sb.Assign(room.FieldName, request.Name))
	}

	query, args := sb.Update(room.Table).Set(sets...).Where(sb.Equal(room.FieldID, filter.ID)).Build()

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

	return r.Get(ctx, &entities.Room{ID: filter.ID})
}

func (r *roomRepo) Get(ctx *gofr.Context, filter *entities.Room) (*entities.Room, error) {
	sb := sqlbuilder.NewSelectBuilder()
	query, args := sb.Select(
		room.FieldID, room.FieldName, room.FieldIsGroup, room.FieldCreatedAt, room.FieldModifiedAt, room.FieldDeletedAt).
		From(room.Table).Where(sb.Equal(room.FieldID, filter.ID)).
		Build()

	row := ctx.SQL.QueryRowContext(ctx, query, args...)
	if err := row.Err(); err != nil {
		return nil, err
	}

	rm := &entities.Room{}

	err := row.Scan(&rm.ID, &rm.Name, &rm.IsGroup, &rm.CreatedAt, &rm.ModifiedAt, &rm.DeletedAt)
	if err != nil {
		return nil, err
	}

	members, err := r.getRoomMembers(ctx, []string{rm.ID})
	if err != nil {
		return nil, err
	}

	rm.Members = members[rm.ID]

	return rm, nil
}

func (r *roomRepo) List(ctx *gofr.Context, filter *repos.RoomFilter) ([]*entities.Room, error) {
	sb := sqlbuilder.NewSelectBuilder().Select(
		room.FieldID, room.FieldName, room.FieldIsGroup, room.FieldCreatedAt, room.FieldModifiedAt, room.FieldDeletedAt).
		From(room.Table)

	if filter.UserID != "" {
		sbMember := sqlbuilder.NewSelectBuilder().
			Select(roommember.FieldRoomID).
			From(roommember.Table)
		sbMember = sbMember.Where(sbMember.EQ(roommember.FieldUserID, filter.UserID))

		sb = sb.Where(sb.In(room.FieldID, sbMember))
	}

	size := 20
	if filter.Size != 0 {
		size = filter.Size
	}

	if filter.Page != 0 {
		sb = sb.Offset((filter.Page - 1) * size)
	}

	query, args := sb.Limit(size).OrderBy(room.FieldModifiedAt).Desc().Build()

	rows, err := ctx.SQL.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rooms := make([]*entities.Room, 0, size)

	for rows.Next() {
		rm := &entities.Room{}
		err = rows.Scan(&rm.ID, &rm.Name, &rm.IsGroup, &rm.CreatedAt, &rm.ModifiedAt, &rm.DeletedAt)
		if err != nil {
			return nil, err
		}

		rooms = append(rooms, rm)
	}

	if slices.Contains(filter.Include, roommember.Entity) {
		roomIds := make([]string, len(rooms))

		for i, rm := range rooms {
			roomIds[i] = rm.ID
		}

		membersMap, err := r.getRoomMembers(ctx, roomIds)
		if err != nil {
			return nil, err
		}

		for _, rm := range rooms {
			rm.Members = membersMap[rm.ID]
		}
	}

	return rooms, nil
}

func (r *roomRepo) Delete(ctx *gofr.Context, filter *entities.Room) error {
	sb := sqlbuilder.NewDeleteBuilder()
	query, args := sb.DeleteFrom(room.Table).Where(sb.Equal(room.FieldID, filter.ID)).Build()

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

func (r *roomRepo) getRoomMembers(ctx *gofr.Context, roomIDs []string) (map[string][]*entities.RoomMember, error) {
	ids := make([]interface{}, len(roomIDs))
	for i, id := range roomIDs {
		ids[i] = id
	}

	sb := sqlbuilder.NewSelectBuilder()
	query, args := sb.
		Select(roommember.FieldRoomID, roommember.FieldUserID, roommember.FieldRole, roommember.FieldCreatedAt).
		From(roommember.Table).
		Where(sb.In(roommember.FieldRoomID, ids...)).
		Build()

	rows, err := ctx.SQL.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	membersMap := make(map[string][]*entities.RoomMember, len(roomIDs))
	for _, roomID := range roomIDs {
		membersMap[roomID] = []*entities.RoomMember{}
	}

	for rows.Next() {
		member := &entities.RoomMember{}

		err = rows.Scan(&member.RoomID, &member.UserID, &member.Role, &member.CreatedAt)
		if err != nil {
			return nil, err
		}

		membersMap[member.RoomID] = append(membersMap[member.RoomID], member)
	}

	return membersMap, nil
}
