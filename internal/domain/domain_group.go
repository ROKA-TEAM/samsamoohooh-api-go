package domain

import (
	"context"
	"time"
)

type Group struct {
	ID          int
	BookTitle   string
	Author      string
	MaxPage     int
	Publisher   string
	Description string
	Bookmark    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time

	// relation
	Users  []*User
	Tasks  []*Task
	Groups []*Group
}

type GroupRepository interface {
	Create(ctx context.Context, group *Group) (*Group, error)
	List(ctx context.Context, offset, limit int) ([]*Group, error)
	GetByID(ctx context.Context, id int) (*Group, error)
	GetUsersByID(ctx context.Context, id int, offset, limit int) ([]*User, error)
	GetPostsByID(ctx context.Context, id int, offset, limit int) ([]*Post, error)
	GetTasksByID(ctx context.Context, id int, offset, limit int) ([]*Task, error)
	Update(ctx context.Context, id int, group *Group) (*Group, error)
	Delete(ctx context.Context, id int) error
}

type GroupService interface {
	Create(ctx context.Context, group *Group) (*Group, error)
	List(ctx context.Context, offset, limit int) ([]*Group, error)
	GetByID(ctx context.Context, id int) (*Group, error)
	GetUsersByID(ctx context.Context, id int, offset, limit int) ([]*User, error)
	GetPostsByID(ctx context.Context, id int, offset, limit int) ([]*Post, error)
	GetTasksByID(ctx context.Context, id int, offset, limit int) ([]*Task, error)
	Update(ctx context.Context, id int, group *Group) (*Group, error)
	Delete(ctx context.Context, id int) error
}
