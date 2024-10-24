package port

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
)

type PostRepository interface {
	CreatePost(ctx context.Context, groupID int, post *domain.Post) (*domain.Post, error)
	GetPosts(ctx context.Context, offset, limit int) ([]*domain.Post, error)
	GetByPostID(ctx context.Context, id int) (*domain.Post, error)
	GetCommentsByPostID(ctx context.Context, id, offset, limit int) ([]*domain.Comment, error)
	UpdatePost(ctx context.Context, id int, post *domain.Post) (*domain.Post, error)
	DeletePost(ctx context.Context, id int) error
}

type PostService interface {
	CreatePost(ctx context.Context, groupID int, post *domain.Post) (*domain.Post, error)
	GetPosts(ctx context.Context, offset, limit int) ([]*domain.Post, error)
	GetByPostID(ctx context.Context, id int) (*domain.Post, error)
	GetCommentsByPostID(ctx context.Context, id, offset, limit int) ([]*domain.Comment, error)
	UpdatePost(ctx context.Context, id int, post *domain.Post) (*domain.Post, error)
	DeletePost(ctx context.Context, id int) error
}
