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
