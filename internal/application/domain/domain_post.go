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
	CreatePost(ctx context.Context, groupID int, post *Post) (*Post, error)
	GetPosts(ctx context.Context, offset, limit int) ([]*Post, error)
	GetByPostID(ctx context.Context, id int) (*Post, error)
	GetCommentsByPostID(ctx context.Context, id, offset, limit int) ([]*Comment, error)
	UpdatePost(ctx context.Context, id int, post *Post) (*Post, error)
	DeletePost(ctx context.Context, id int) error
}

type PostService interface {
	CreatePost(ctx context.Context, groupID int, post *Post) (*Post, error)
	GetPosts(ctx context.Context, offset, limit int) ([]*Post, error)
	GetByPostID(ctx context.Context, id int) (*Post, error)
	GetCommentsByPostID(ctx context.Context, id, offset, limit int) ([]*Comment, error)
	UpdatePost(ctx context.Context, id int, post *Post) (*Post, error)
	DeletePost(ctx context.Context, id int) error
}
