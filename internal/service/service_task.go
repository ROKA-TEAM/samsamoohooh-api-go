package service

import (
	"context"
	"samsamoohooh-go-api/internal/domain"
)

var _ domain.TaskService = (*TaskService)(nil)

type TaskService struct {
	taskRepository domain.TaskRepository
}

func (s *TaskService) Create(ctx context.Context, groupID int, task *domain.Task) (*domain.Task, error) {
	return s.taskRepository.Create(ctx, groupID, task)
}

func (s *TaskService) List(ctx context.Context, id, offset, limit int) ([]*domain.Task, error) {
	return s.taskRepository.List(ctx, id, offset, limit)
}

func (s *TaskService) GetByID(ctx context.Context, id int) (*domain.Task, error) {
	return s.taskRepository.GetByID(ctx, id)
}

func (s *TaskService) GetTopicsByID(ctx context.Context, id, offset, limit int) ([]*domain.Topic, error) {
	return s.taskRepository.GetTopicsByID(ctx, id, offset, limit)
}

func (s *TaskService) Updated(ctx context.Context, id int, task *domain.Task) (*domain.Task, error) {
	return s.taskRepository.Updated(ctx, id, task)
}

func (s *TaskService) Delete(ctx context.Context, id int) error {
	return s.taskRepository.Delete(ctx, id)
}
