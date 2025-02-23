package services

import (
	"time"
)

type ListSpaces struct {
	UserID        string
	UpdatedBefore time.Time
}

type CreateMessage struct {
	UserID     string
	Content    string
	Recipients []string
}

type ListMessages struct {
	UserID        string
	SpaceID       string
	CreatedBefore time.Time
}
