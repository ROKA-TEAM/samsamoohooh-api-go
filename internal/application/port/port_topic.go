package port

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
)

type TopicRepository interface {
	CreateTopic(ctx context.Context, taskID int, topic *domain.Topic) (*domain.Topic, error)
	GetTopics(ctx context.Context, offset, limit int) ([]*domain.Topic, error)
	GetByTopicID(ctx context.Context, id int) (*domain.Topic, error)
	UpdateTopic(ctx context.Context, id int, topic *domain.Topic) (*domain.Topic, error)
	DeleteTopic(ctx context.Context, id int) error
}

type TopicService interface {
	CreateTopic(ctx context.Context, taskID int, topic *domain.Topic) (*domain.Topic, error)
	GetTopics(ctx context.Context, offset, limit int) ([]*domain.Topic, error)
	GetByTopicID(ctx context.Context, id int) (*domain.Topic, error)
	UpdateTopic(ctx context.Context, id int, topic *domain.Topic) (*domain.Topic, error)
	DeleteTopic(ctx context.Context, id int) error
}
