package roommember

const (
	Entity         string = "members"
	Table          string = "room_members"
	FieldRoomID    string = "room_id"
	FieldUserID    string = "user_id"
	FieldRole      string = "role"
	FieldCreatedAt string = "created_at"
)

type Role string

const (
	RoleAdmin  Role = "ADMIN"
	RoleMember Role = "MEMBER"
)
