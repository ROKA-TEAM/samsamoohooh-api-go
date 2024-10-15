package domain

import "time"

type Group struct {
	ID          int
	BookTitle   string
	Author      string
	MaxPage     int
	Publisher   string
	Description string
	Bookmark    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time

	// relation
	Users  []*User
	Tasks  []*Task
	Groups []*Group
}
