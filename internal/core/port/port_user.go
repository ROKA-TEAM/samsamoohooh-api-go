package port

import (
	"context"
	"samsamoohooh-go-api/internal/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	GetByID(ctx context.Context, id uint) (*domain.User, error)
	// 반환이 []*Domain.Group인 이유는 many 2 many 관계를 정의할 때 포인터를 사용해 정의했기 때문에
	GetGroupsByID(ctx context.Context, id uint) ([]*domain.Group, error)
	// GetPostsByID(ctx context.Context, id uint) ([]domain.Post, error)
	// GetCommentsByID(ctx context.Context, id uint) ([]domain.Comment, error)
	// GetSubjectsByID(ctx context.Context, id uint) ([]domain.Subject, error)
	GetAll(ctx context.Context, skip, limit int) ([]domain.User, error)
	Update(ctx context.Context, id uint, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id uint) error
}
