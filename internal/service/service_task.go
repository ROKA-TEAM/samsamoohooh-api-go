package service

import (
	"context"
	domain2 "samsamoohooh-go-api/internal/application/domain"
)

var _ domain2.TaskService = (*TaskService)(nil)

type TaskService struct {
	taskRepository domain2.TaskRepository
}

func NewTaskService(taskRepository domain2.TaskRepository) *TaskService {
	return &TaskService{taskRepository: taskRepository}
}

func (s *TaskService) Create(ctx context.Context, groupID int, task *domain2.Task) (*domain2.Task, error) {
	return s.taskRepository.Create(ctx, groupID, task)
}

func (s *TaskService) List(ctx context.Context, offset, limit int) ([]*domain2.Task, error) {
	return s.taskRepository.List(ctx, offset, limit)
}

func (s *TaskService) GetByID(ctx context.Context, id int) (*domain2.Task, error) {
	return s.taskRepository.GetByID(ctx, id)
}

func (s *TaskService) GetTopicsByID(ctx context.Context, id, offset, limit int) ([]*domain2.Topic, error) {
	return s.taskRepository.GetTopicsByID(ctx, id, offset, limit)
}

func (s *TaskService) Updated(ctx context.Context, id int, task *domain2.Task) (*domain2.Task, error) {
	return s.taskRepository.Updated(ctx, id, task)
}

func (s *TaskService) Delete(ctx context.Context, id int) error {
	return s.taskRepository.Delete(ctx, id)
}
