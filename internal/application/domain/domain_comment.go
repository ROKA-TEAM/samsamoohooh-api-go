package domain

import (
	"time"
)

type Comment struct {
	ID        int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	// relation
	UserID int
	PostID int
}
