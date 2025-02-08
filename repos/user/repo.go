package user

import (
	"database/sql"
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/entities/user"
	"github.com/hnpatil/gochat/repos"
	"github.com/huandu/go-sqlbuilder"
	"gofr.dev/pkg/gofr"
	"time"
)

type userRepo struct {
}

func New() repos.User {
	return &userRepo{}
}

func (u *userRepo) Create(ctx *gofr.Context, request *entities.User) (*entities.User, error) {
	query, args := sqlbuilder.NewInsertBuilder().
		InsertInto(user.Table).
		Cols(user.FieldID, user.FieldName).
		Values(request.ID, request.Name).
		Build()

	_, err := ctx.SQL.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, repos.Error(err, user.Entity)
	}

	return u.Get(ctx, &entities.User{ID: request.ID})
}

func (u *userRepo) Update(ctx *gofr.Context, filter, request *entities.User) (*entities.User, error) {
	sb := sqlbuilder.NewUpdateBuilder()
	sets := []string{sb.Assign(user.FieldModifiedAt, time.Now())}

	if request.Name != "" {
		sets = append(sets, sb.Assign(user.FieldName, request.Name))
	}

	query, args := sb.Update(user.Table).Set(sets...).Where(sb.Equal(user.FieldID, filter.ID)).Build()

	r, err := ctx.SQL.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, repos.Error(err, user.Entity)
	}

	n, err := r.RowsAffected()
	if err != nil {
		return nil, repos.Error(err, user.Entity)
	}

	if n == 0 {
		return nil, repos.Error(sql.ErrNoRows, user.Entity)
	}

	return u.Get(ctx, &entities.User{ID: filter.ID})
}

func (u *userRepo) Get(ctx *gofr.Context, filter *entities.User) (*entities.User, error) {
	sb := sqlbuilder.NewSelectBuilder()
	query, args := sb.Select(
		user.FieldID, user.FieldName, user.FieldCreatedAt, user.FieldModifiedAt, user.FieldDeletedAt).
		From(user.Table).Where(sb.Equal(user.FieldID, filter.ID)).Build()

	row := ctx.SQL.QueryRowContext(ctx, query, args...)
	if err := row.Err(); err != nil {
		return nil, repos.Error(err, user.Entity)
	}

	usr := &entities.User{}

	err := row.Scan(&usr.ID, &usr.Name, &usr.CreatedAt, &usr.ModifiedAt, &usr.DeletedAt)
	if err != nil {
		return nil, repos.Error(err, user.Entity)
	}

	return usr, nil
}

func (u *userRepo) List(ctx *gofr.Context, filter *repos.UserFilter) ([]*entities.User, error) {
	sb := sqlbuilder.NewSelectBuilder().
		Select(user.FieldID, user.FieldName, user.FieldCreatedAt, user.FieldModifiedAt, user.FieldDeletedAt).
		From(user.Table)

	size := 20
	if filter.Size != 0 {
		size = filter.Size
	}

	if filter.Page != 0 {
		sb = sb.Offset((filter.Page - 1) * size)
	}

	if len(filter.UserID) > 0 {
		ids := make([]interface{}, len(filter.UserID))
		for i, v := range filter.UserID {
			ids[i] = v
		}

		sb = sb.Where(sb.In(user.FieldID, ids...))
	}

	query, args := sb.Limit(size).Build()
	rows, err := ctx.SQL.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, repos.Error(err, user.Entity)
	}

	defer rows.Close()

	users := make([]*entities.User, 0, size)

	for rows.Next() {
		usr := &entities.User{}
		err = rows.Scan(&usr.ID, &usr.Name, &usr.CreatedAt, &usr.ModifiedAt, &usr.DeletedAt)
		if err != nil {
			return nil, repos.Error(err, user.Entity)
		}

		users = append(users, usr)
	}

	return users, repos.Error(err, user.Entity)
}

func (u *userRepo) Delete(ctx *gofr.Context, filter *entities.User) error {
	sb := sqlbuilder.NewDeleteBuilder()
	query, args := sb.DeleteFrom(user.Table).Where(sb.Equal(user.FieldID, filter.ID)).Build()

	r, err := ctx.SQL.ExecContext(ctx, query, args...)
	if err != nil {
		return repos.Error(err, user.Entity)
	}

	n, err := r.RowsAffected()
	if err != nil {
		return repos.Error(err, user.Entity)
	}

	if n == 0 {
		return repos.Error(sql.ErrNoRows, user.Entity)
	}

	return nil
}
