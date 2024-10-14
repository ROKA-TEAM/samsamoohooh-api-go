package domain

import "time"

type Group struct {
	ID          int
	BookTitle   string
	Author      string
	MaxPage     int
	Publisher   string
	Description string
	Bookmark    int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Users []User
	Posts []Post
	Tasks []Task
}
