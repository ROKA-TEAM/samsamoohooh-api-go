package repository

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/repository/database"
	"samsamoohooh-go-api/internal/application/repository/database/utils"
)

var _ domain.TopicRepository = (*TopicRepository)(nil)

type TopicRepository struct {
	database *database.Database
}

func NewTopicRepository(database *database.Database) *TopicRepository {
	return &TopicRepository{database: database}
}

func (r *TopicRepository) CreateTopic(ctx context.Context, taskID int, topic *domain.Topic) (*domain.Topic, error) {
	createdTopic, err := r.database.Topic.
		Create().
		SetTopic(topic.Topic).
		SetFeeling(topic.Feeling).
		SetTaskID(taskID).
		Save(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainTopic(createdTopic), nil
}
func (r *TopicRepository) GetTopics(ctx context.Context, offset, limit int) ([]*domain.Topic, error) {
	listTopic, err := r.database.Topic.
		Query().
		Offset(offset).
		Limit(limit).
		All(ctx)

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainTopics(listTopic), nil
}

func (r *TopicRepository) GetByTopicID(ctx context.Context, id int) (*domain.Topic, error) {
	gotTopic, err := r.database.Topic.
		Get(ctx, id)
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainTopic(gotTopic), nil
}

func (r *TopicRepository) UpdateTopic(ctx context.Context, id int, topic *domain.Topic) (*domain.Topic, error) {
	updateBuilder := r.database.Topic.
		UpdateOneID(id)

	if topic.Topic != "" {
		updateBuilder.SetTopic(topic.Topic)
	}

	if topic.Feeling != "" {
		updateBuilder.SetFeeling(topic.Feeling)
	}

	updatedTopic, err := updateBuilder.Save(ctx)
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return utils.ConvertDomainTopic(updatedTopic), nil
}

func (r *TopicRepository) DeleteTopic(ctx context.Context, id int) error {
	err := r.database.Topic.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return utils.Wrap(err)
	}

	return nil
}
