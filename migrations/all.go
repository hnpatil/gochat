package migrations

import "gofr.dev/pkg/gofr/migration"

const (
	createUsers = `create table users(
      id TEXT PRIMARY KEY,
      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      deleted_at TIMESTAMP,
      name TEXT NOT NULL
	);`

	createRooms = `create table rooms(
      id TEXT PRIMARY KEY,
      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      deleted_at TIMESTAMP,
      name TEXT NOT NULL,
      is_group BOOLEAN NOT NULL DEFAULT FALSE
	);`

	createRoomMembers = `CREATE TABLE room_members (
		room_id TEXT NOT NULL,
		user_id TEXT NOT NULL,
		role TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (room_id, user_id),
		FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
  	);`

	createMessages = `CREATE TABLE messages (
		id UUID PRIMARY KEY,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP,
		room_id TEXT NOT NULL,
		sender_id TEXT NOT NULL,
		sent_at TIMESTAMP,
		status TEXT NOT NULL,
		content TEXT NOT NULL,
		FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE,
		FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE SET NULL
	);`
)

func All() map[int64]migration.Migrate {
	return map[int64]migration.Migrate{
		20250105010000: executeMigration(createUsers),
		20250105030000: executeMigration(createRooms),
		20250105040000: executeMigration(createRoomMembers),
		20250105050000: executeMigration(createMessages),
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
