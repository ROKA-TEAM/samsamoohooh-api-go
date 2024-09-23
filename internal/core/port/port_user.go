package port

import (
	"context"
	"samsamoohooh-go-api/internal/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)

	GetByID(ctx context.Context, id uint) (*domain.User, error)
	GetGroupsByID(ctx context.Context, id int) ([]*domain.Group, error)
	GetAll(ctx context.Context, skip, limit int) ([]*domain.User, error)

	Update(ctx context.Context, id uint, user *domain.User) (*domain.User, error)

	Delete(ctx context.Context, id uint) error
}
