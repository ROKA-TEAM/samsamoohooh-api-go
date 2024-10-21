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
	CreateTopic(ctx context.Context, taskID int, topic *Topic) (*Topic, error)
	GetTopics(ctx context.Context, offset, limit int) ([]*Topic, error)
	GetByTopicID(ctx context.Context, id int) (*Topic, error)
	UpdateTopic(ctx context.Context, id int, topic *Topic) (*Topic, error)
	DeleteTopic(ctx context.Context, id int) error
}

type TopicService interface {
	CreateTopic(ctx context.Context, taskID int, topic *Topic) (*Topic, error)
	GetTopics(ctx context.Context, offset, limit int) ([]*Topic, error)
	GetByTopicID(ctx context.Context, id int) (*Topic, error)
	UpdateTopic(ctx context.Context, id int, topic *Topic) (*Topic, error)
	DeleteTopic(ctx context.Context, id int) error
}
