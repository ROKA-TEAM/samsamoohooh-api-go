package service

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
)

var _ domain.TaskService = (*TaskService)(nil)

type TaskService struct {
	taskRepository domain.TaskRepository
}

func NewTaskService(
	taskRepository domain.TaskRepository,
) *TaskService {
	return &TaskService{taskRepository: taskRepository}
}

func (s *TaskService) CreateTask(ctx context.Context, groupID int, task *domain.Task) (*domain.Task, error) {
	createdTask, err := s.taskRepository.CreateTask(ctx, groupID, task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (s *TaskService) GetTasks(ctx context.Context, offset, limit int) ([]*domain.Task, error) {
	listTask, err := s.taskRepository.GetTasks(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	return listTask, nil
}

func (s *TaskService) GetByTaskID(ctx context.Context, id int) (*domain.Task, error) {
	gotTask, err := s.taskRepository.GetByTaskID(ctx, id)
	if err != nil {
		return nil, err
	}

	return gotTask, nil
}

func (s *TaskService) GetTopicsByTaskID(ctx context.Context, id, offset, limit int) ([]*domain.Topic, error) {
	listTopic, err := s.taskRepository.GetTopicsByTaskID(ctx, id, offset, limit)
	if err != nil {
		return nil, err
	}

	return listTopic, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id int, task *domain.Task) (*domain.Task, error) {
	updatedTask, err := s.taskRepository.UpdateTask(ctx, id, task)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	err := s.taskRepository.DeleteTask(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) GetTopicsLenByTaskID(ctx context.Context, id int) (int, error) {
	topicsLen, err := s.taskRepository.GetTopicsLenByTaskID(ctx, id)
	if err != nil {
		return 0, err
	}

	return topicsLen, nil
}
