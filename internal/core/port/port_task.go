package port

import (
	"context"
	"samsamoohooh-go-api/internal/core/domain"
)

type TaskRepository interface {
	Create(ctx context.Context, task *domain.Task) (*domain.Task, error)
	GetByID(ctx context.Context, id uint) (*domain.Task, error)
	GetSubjectsByID(ctx context.Context, id uint) ([]domain.Subject, error)
	GetAll(ctx context.Context, skip, limit int) ([]domain.Task, error)
	Update(ctx context.Context, id uint, task *domain.Task) (*domain.Task, error)
	Delete(ctx context.Context, id uint) error
}
