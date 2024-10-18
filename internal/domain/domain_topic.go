package domain

import (
	"context"
	"time"
)

type Topic struct {
	ID      int
	Topic   string
	Feeling string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	// relation
	UserID int
	TaskID int
}

type TopicRepository interface {
	Create(ctx context.Context, topic *Topic) (*Topic, error)
	List(ctx context.Context, offset, limit int) ([]*Topic, error)
	GetByID(ctx context.Context, id int) (Topic, error)
	Update(ctx context.Context, id int, topic *Topic) (Topic, error)
	Delete(ctx context.Context, id int) error
}

type TopicService interface {
	Create(ctx context.Context, topic *Topic) (*Topic, error)
	List(ctx context.Context, offset, limit int) ([]*Topic, error)
	GetByID(ctx context.Context, id int) (Topic, error)
	Update(ctx context.Context, id int, topic *Topic) (Topic, error)
	Delete(ctx context.Context, id int) error
}
