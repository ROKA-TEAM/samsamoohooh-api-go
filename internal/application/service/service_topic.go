package service

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
)

var _ domain.TopicService = (*TopicService)(nil)

type TopicService struct {
	topicRepository domain.TopicRepository
}

func NewTopicService(
	topicRepository domain.TopicRepository,
) *TopicService {
	return &TopicService{topicRepository: topicRepository}
}

func (s *TopicService) CreateTopic(ctx context.Context, taskID int, topic *domain.Topic) (*domain.Topic, error) {
	createdTopic, err := s.topicRepository.CreateTopic(ctx, taskID, topic)
	if err != nil {
		return nil, err
	}

	return createdTopic, nil
}

func (s *TopicService) GetTopics(ctx context.Context, offset, limit int) ([]*domain.Topic, error) {
	listTopic, err := s.topicRepository.GetTopics(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	return listTopic, nil
}
func (s *TopicService) GetByTopicID(ctx context.Context, id int) (*domain.Topic, error) {
	gotTopic, err := s.topicRepository.GetByTopicID(ctx, id)
	if err != nil {
		return nil, err
	}

	return gotTopic, nil
}

func (s *TopicService) UpdateTopic(ctx context.Context, id int, topic *domain.Topic) (*domain.Topic, error) {
	updatedTopic, err := s.topicRepository.UpdateTopic(ctx, id, topic)
	if err != nil {
		return nil, err
	}

	return updatedTopic, nil
}

func (s *TopicService) DeleteTopic(ctx context.Context, id int) error {

	err := s.topicRepository.DeleteTopic(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
