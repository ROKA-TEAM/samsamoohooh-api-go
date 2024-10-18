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
	Create(ctx context.Context, comment *Comment) (*Comment, error)
	List(ctx context.Context, offset, limit int) ([]*Comment, error)
	GetByID(ctx context.Context, id int) (*Comment, error)
	Update(ctx context.Context, id int, comment *Comment) (*Comment, error)
	Delete(ctx context.Context, id int) error
}

type CommentService interface {
	Create(ctx context.Context, comment *Comment) (*Comment, error)
	List(ctx context.Context, offset, limit int) ([]*Comment, error)
	GetByID(ctx context.Context, id int) (*Comment, error)
	Update(ctx context.Context, id int, comment *Comment) (*Comment, error)
	Delete(ctx context.Context, id int) error
}
