package repos

import "time"

type MessageFilter struct {
	CreatedBefore time.Time
	SpaceID       string
}

type SpaceFilter struct {
	UserID        string
	UpdatedBefore time.Time
}
