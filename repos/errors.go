package repos

import (
	"database/sql"
	"errors"
	chatErrors "github.com/hnpatil/gochat/errors"
	"github.com/lib/pq"
)

func Error(err error, entity string) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return chatErrors.EntityNotFound(entity)
	}

	pErr := &pq.Error{}
	if errors.As(err, &pErr) {
		if pErr.Code == "23505" {
			return chatErrors.EntityExists(entity)
		}
	}

	return err
}
