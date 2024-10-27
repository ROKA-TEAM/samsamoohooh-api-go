package domain

import (
	"time"
)

type Topic struct {
	ID      int
	Topic   string
	Feeling string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	// relation
	UserID int
	TaskID int
}
