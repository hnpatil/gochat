package migrations

import (
	"github.com/gocql/gocql"
	"gofr.dev/pkg/gofr"
)

func Run(app *gofr.App, cass *gocql.Session) error {
	app.Logger().Info("Running migrations")

	err := cass.Query(createMessages).Exec()
	if err != nil {
		return err
	}

	err = cass.Query(createUserSpaces).Exec()
	if err != nil {
		return err
	}

	err = cass.Query(createUserSpacesView).Exec()
	if err != nil {
		return err
	}

	return nil
}
