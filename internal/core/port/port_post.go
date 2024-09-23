package port

import (
	"context"
	"samsamoohooh-go-api/internal/core/domain"
)

type PostRepository interface {
	Create(ctx context.Context, post *domain.Post) (*domain.Post, error)
	GetByID(ctx context.Context, id uint) (*domain.Post, error)
	GetCommentsByID(ctx context.Context, id uint) ([]*domain.Comment, error)
	GetAll(ctx context.Context, skip, limit int) ([]*domain.Post, error)
	Update(ctx context.Context, id uint, post *domain.Post) (*domain.Post, error)
	Delete(ctx context.Context, id uint) error
}
