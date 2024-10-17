package domain

import (
	"context"
	"time"
)

type Post struct {
	ID      int
	Title   string
	Content string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	// relation
	Comments []*Comment

	UserID  int
	GroupID int
}

type PostRepository interface {
	Create(ctx context.Context, post *Post) (*Post, error)
	List(ctx context.Context, offset, limit int) ([]*Post, error)
	GetByID(ctx context.Context, id int) (*Post, error)
	GetCommentsByID(ctx context.Context, id int) ([]*Comment, error)
	Update(ctx context.Context, id int, post *Post) (*Post, error)
	Delete(ctx context.Context, id int) error
}

type PostService interface {
	Create(ctx context.Context, post *Post) (*Post, error)
	List(ctx context.Context, offset, limit int) ([]*Post, error)
	GetByID(ctx context.Context, id int) (*Post, error)
	GetCommentsByID(ctx context.Context, id int) ([]*Comment, error)
	Update(ctx context.Context, id int, post *Post) (*Post, error)
	Delete(ctx context.Context, id int) error
}
