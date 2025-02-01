package services

type CreateUser struct {
	ID   string
	Name string
}

type UpdateUser struct {
	ID   string
	Name string
}

type GetUser struct {
	ID string
}

type DeleteUser struct {
	ID string
}

type ListUsers struct {
	Page   int
	Size   int
	UserID string
}

type CreateRoom struct {
	RoomID  string
	UserID  string
	Name    string
	Members []string
}

type UpdateRoom struct {
	UserID string
	RoomID string
	Name   string
}

type GetRoom struct {
	UserID string
	RoomID string
}

type ListRoom struct {
	Include []string
	UserID  string
	Page    int
	Size    int
}

type DeleteRoom struct {
	RoomID string
	UserID string
}
