package repository

import (
	"context"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/database"
	"samsamoohooh-go-api/internal/core/domain"
	"samsamoohooh-go-api/internal/core/port"
)

var _ port.TaskRepository = (*TaskRepository)(nil)

type TaskRepository struct {
	database *database.Database
}

func NewTaskRepository(database *database.Database) *TaskRepository {
	return &TaskRepository{
		database: database,
	}
}
func (r *TaskRepository) Create(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	err := r.database.WithContext(ctx).Create(task).Error
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id uint) (*domain.Task, error) {
	task := domain.Task{}
	err := r.database.WithContext(ctx).First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}
func (r *TaskRepository) GetSubjectsByID(ctx context.Context, id uint) ([]domain.Subject, error) {
	task := domain.Task{}
	err := r.database.WithContext(ctx).Preload("Subjects").First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return task.Subjects, nil
}

func (r *TaskRepository) GetAll(ctx context.Context, skip, limit int) ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.database.WithContext(ctx).Limit(limit).Offset((skip - 1) * limit).Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
func (r *TaskRepository) Update(ctx context.Context, id uint, task *domain.Task) (*domain.Task, error) {
	task.ID = id
	err := r.database.WithContext(ctx).Save(task).Error
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) Delete(ctx context.Context, id uint) error {
	err := r.database.WithContext(ctx).Delete(&domain.Task{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
