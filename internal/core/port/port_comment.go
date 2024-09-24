package port

import (
	"context"
	"samsamoohooh-go-api/internal/core/domain"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *domain.Comment) (*domain.Comment, error)
	GetByID(ctx context.Context, id uint) (*domain.Comment, error)
	GetAll(ctx context.Context, skip, limit int) ([]domain.Comment, error)
	Update(ctx context.Context, id uint, comment *domain.Comment) (*domain.Comment, error)
	Delete(ctx context.Context, id uint) error
}
