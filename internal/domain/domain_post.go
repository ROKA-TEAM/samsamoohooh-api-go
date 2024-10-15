package domain

import "time"

type Post struct {
	ID      int
	Title   string
	Content string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	// relation
	Comments []*Comment

	UserID  int
	GroupID int
}
