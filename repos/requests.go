package repos

type UserFilter struct {
	Page   int
	Size   int
	UserID string
}

type RoomFilter struct {
	Page    int
	Size    int
	UserID  string
	Include []string
}
