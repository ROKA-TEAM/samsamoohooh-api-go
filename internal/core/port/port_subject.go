package port

import (
	"context"
	"samsamoohooh-go-api/internal/core/domain"
)

type SubjectRepository interface {
	Create(ctx context.Context, subject *domain.Subject) (*domain.Subject, error)
	GetByID(ctx context.Context, id uint) (*domain.Subject, error)
	GetAll(ctx context.Context, skip, limit int) ([]domain.Subject, error)
	Update(ctx context.Context, id uint, subject *domain.Subject) (*domain.Subject, error)
	Delete(ctx context.Context, id uint) error
}
