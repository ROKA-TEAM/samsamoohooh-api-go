package domain

import (
	"context"
	"time"
)

type Comment struct {
	ID        int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	// relation
	UserID int
	PostID int
}

type CommentRepository interface {
	CreateComment(ctx context.Context, postID int, comment *Comment) (*Comment, error)
	GetComments(ctx context.Context, offset, limit int) ([]*Comment, error)
	GetByCommentID(ctx context.Context, id int) (*Comment, error)
	UpdateComment(ctx context.Context, id int, comment *Comment) (*Comment, error)
	DeleteComment(ctx context.Context, id int) error
}

type CommentService interface {
	CreateComment(ctx context.Context, postID int, comment *Comment) (*Comment, error)
	GetComments(ctx context.Context, offset, limit int) ([]*Comment, error)
	GetByCommentID(ctx context.Context, id int) (*Comment, error)
	UpdateComment(ctx context.Context, id int, comment *Comment) (*Comment, error)
	DeleteComment(ctx context.Context, id int) error
}
