package port

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, postID int, comment *domain.Comment) (*domain.Comment, error)
	GetComments(ctx context.Context, offset, limit int) ([]*domain.Comment, error)
	GetByCommentID(ctx context.Context, id int) (*domain.Comment, error)
	UpdateComment(ctx context.Context, id int, comment *domain.Comment) (*domain.Comment, error)
	DeleteComment(ctx context.Context, id int) error
}

type CommentService interface {
	CreateComment(ctx context.Context, postID int, comment *domain.Comment) (*domain.Comment, error)
	GetComments(ctx context.Context, offset, limit int) ([]*domain.Comment, error)
	GetByCommentID(ctx context.Context, id int) (*domain.Comment, error)
	UpdateComment(ctx context.Context, id int, comment *domain.Comment) (*domain.Comment, error)
	DeleteComment(ctx context.Context, id int) error
}
