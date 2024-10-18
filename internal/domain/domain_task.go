package domain

import (
	"context"
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

type TaskRepository interface {
	Create(ctx context.Context, groupID int, task *Task) (*Task, error)
	List(ctx context.Context, id, offset, limit int) ([]*Task, error)
	GetByID(ctx context.Context, id int) (*Task, error)
	GetTopicsByID(ctx context.Context, id, offset, limit int) ([]*Topic, error)
	Updated(ctx context.Context, id int, task *Task) (*Task, error)
	Delete(ctx context.Context, id int) error
}

type TaskService interface {
	Create(ctx context.Context, groupID int, task *Task) (*Task, error)
	List(ctx context.Context, id, offset, limit int) ([]*Task, error)
	GetByID(ctx context.Context, id int) (*Task, error)
	GetTopicsByID(ctx context.Context, id, offset, limit int) ([]*Topic, error)
	Updated(ctx context.Context, id int, task *Task) (*Task, error)
	Delete(ctx context.Context, id int) error
}
