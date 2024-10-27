package domain

import (
	"time"
)

type Task struct {
	ID       int
	Deadline time.Time
	Range    int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	// relation
	Topics []*Topic

	GroupID int
}
