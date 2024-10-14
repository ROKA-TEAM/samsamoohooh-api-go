package domain

import "time"

type Subject struct {
	ID      int
	Content string
	Feeling string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	UserID int
	TaskID int
}
