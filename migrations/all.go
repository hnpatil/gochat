package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

const (
	createMessages = `CREATE TABLE IF NOT EXISTS messages (
      space_id TEXT,
      created_at TIMESTAMP,
      data TEXT,
      PRIMARY KEY (space_id, created_at)
	) WITH CLUSTERING ORDER BY (created_at DESC);`

	createUserSpaces = `CREATE TABLE IF NOT EXISTS user_spaces (
     user_id TEXT,
     space_id TEXT,
     updated_at TIMESTAMP,
     data TEXT,
     PRIMARY KEY (user_id, space_id)
	);`

	createUserSpacesView = `CREATE MATERIALIZED VIEW IF NOT EXISTS user_spaces_view AS
	SELECT user_id, space_id, updated_at, data
	FROM user_spaces
	WHERE user_id IS NOT NULL AND space_id IS NOT NULL AND updated_at IS NOT NULL
	PRIMARY KEY (user_id, updated_at, space_id)
	WITH CLUSTERING ORDER BY (updated_at DESC);`
)

func All() map[int64]migration.Migrate {
	return map[int64]migration.Migrate{
		20250105010000: executeMigration(createMessages),
		20250105020000: executeMigration(createUserSpaces),
		20250105030000: executeMigration(createUserSpacesView),
	}
}

func executeMigration(statement string) migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(statement)
			return err
		},
	}
}
