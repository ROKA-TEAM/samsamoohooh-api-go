package domain

import "time"

type Post struct {
	ID      int
	Title   string
	Content string

	UserID  int
	GroupID int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Comments []Comment
}
