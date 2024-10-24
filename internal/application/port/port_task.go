package port

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, groupID int, task *domain.Task) (*domain.Task, error)
	GetTasks(ctx context.Context, offset, limit int) ([]*domain.Task, error)
	GetByTaskID(ctx context.Context, id int) (*domain.Task, error)
	GetTopicsByTaskID(ctx context.Context, id, offset, limit int) ([]*domain.Topic, error)
	UpdateTask(ctx context.Context, id int, task *domain.Task) (*domain.Task, error)
	DeleteTask(ctx context.Context, id int) error
	GetTopicsLenByTaskID(ctx context.Context, id int) (int, error)
}

type TaskService interface {
	CreateTask(ctx context.Context, groupID int, task *domain.Task) (*domain.Task, error)
	GetTasks(ctx context.Context, offset, limit int) ([]*domain.Task, error)
	GetByTaskID(ctx context.Context, id int) (*domain.Task, error)
	GetTopicsByTaskID(ctx context.Context, id, offset, limit int) ([]*domain.Topic, error)
	UpdateTask(ctx context.Context, id int, task *domain.Task) (*domain.Task, error)
	DeleteTask(ctx context.Context, id int) error
	GetTopicsLenByTaskID(ctx context.Context, id int) (int, error)
}
