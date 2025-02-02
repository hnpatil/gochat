package repos

import "time"

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

type MessageFilter struct {
	ModifiedBefore time.Time
	RoomID         string
}
