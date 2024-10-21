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
	CreateTask(ctx context.Context, groupID int, task *Task) (*Task, error)
	GetTasks(ctx context.Context, offset, limit int) ([]*Task, error)
	GetByTaskID(ctx context.Context, id int) (*Task, error)
	GetTopicsByTaskID(ctx context.Context, id, offset, limit int) ([]*Topic, error)
	UpdateTask(ctx context.Context, id int, task *Task) (*Task, error)
	DeleteTask(ctx context.Context, id int) error
}

type TaskService interface {
	CreateTask(ctx context.Context, groupID int, task *Task) (*Task, error)
	GetTasks(ctx context.Context, offset, limit int) ([]*Task, error)
	GetByTaskID(ctx context.Context, id int) (*Task, error)
	GetTopicsByTaskID(ctx context.Context, id, offset, limit int) ([]*Topic, error)
	UpdateTask(ctx context.Context, id int, task *Task) (*Task, error)
	DeleteTask(ctx context.Context, id int) error
}
