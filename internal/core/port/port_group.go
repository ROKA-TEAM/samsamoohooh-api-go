package port

import (
	"context"
	"samsamoohooh-go-api/internal/core/domain"
)

type GroupRepository interface {
	Create(ctx context.Context, group *domain.Group) (*domain.Group, error)

	GetByID(ctx context.Context, id uint) (*domain.Group, error)
	GetUsersByID(ctx context.Context, id uint) ([]*domain.User, error)
	GetAll(ctx context.Context, skip, limit int) ([]*domain.Group, error)

	Update(ctx context.Context, id uint, group *domain.Group) (*domain.Group, error)

	Delete(ctx context.Context, id uint) error
}
