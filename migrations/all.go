package migrations

import "gofr.dev/pkg/gofr/migration"

const (
	createUsers = `create table users(
      id VARCHAR(36) PRIMARY KEY,
      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      metadata jsonb DEFAULT '{}'
	);`

	createRooms = `create table rooms(
      id VARCHAR(36) PRIMARY KEY,
      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      metadata jsonb DEFAULT '{}'
	);`

	createRoomMembers = `CREATE TABLE room_members (
		room_id VARCHAR(36) NOT NULL,
		user_id VARCHAR(36) NOT NULL,
		role VARCHAR(12) NOT NULL CHECK (role IN ('ADMIN', 'MEMBER')),
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (room_id, user_id),
		FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
  	);`

	createMessages = `CREATE TABLE messages (
		id VARCHAR(8),
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		room_id VARCHAR(36) NOT NULL,
		sender_id VARCHAR(36) NOT NULL,
		sent_at TIMESTAMP,
		content TEXT NOT NULL,
		PRIMARY KEY (room_id, id),
		FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE,
		FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE SET NULL
	);`

	createModifiedAtIndexOnMessages = "CREATE INDEX idx_messages_modified_at ON messages(modified_at);"
	createModifiedAtIndexOnRooms    = "CREATE INDEX idx_rooms_modified_at ON rooms(modified_at);"
)

func All() map[int64]migration.Migrate {
	return map[int64]migration.Migrate{
		20250105010000: executeMigration(createUsers),
		20250105030000: executeMigration(createRooms),
		20250105040000: executeMigration(createRoomMembers),
		20250105050000: executeMigration(createMessages),
		20250105060000: executeMigration(createModifiedAtIndexOnMessages),
		20250105070000: executeMigration(createModifiedAtIndexOnRooms),
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
