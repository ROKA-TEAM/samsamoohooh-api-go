package service

import (
	"context"
	"samsamoohooh-go-api/internal/domain"
)

type TopicService struct {
	topicRepository domain.TopicRepository
}

func NewTopicService(topicRepository domain.TopicRepository) *TopicService {
	return &TopicService{topicRepository: topicRepository}
}

func (s *TopicService) Create(ctx context.Context, taskID int, topic *domain.Topic) (*domain.Topic, error) {
	return s.topicRepository.Create(ctx, taskID, topic)
}
func (s *TopicService) List(ctx context.Context, offset, limit int) ([]*domain.Topic, error) {
	return s.topicRepository.List(ctx, offset, limit)
}
func (s *TopicService) GetByID(ctx context.Context, id int) (*domain.Topic, error) {
	return s.topicRepository.GetByID(ctx, id)
}
func (s *TopicService) Update(ctx context.Context, id int, topic *domain.Topic) (*domain.Topic, error) {
	return s.topicRepository.Update(ctx, id, topic)
}
func (s *TopicService) Delete(ctx context.Context, id int) error {
	return s.topicRepository.Delete(ctx, id)
}
