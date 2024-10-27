package presenter

import (
	"samsamoohooh-go-api/internal/application/domain"
)

type TopicCreateResponse struct {
	ID      int    `json:"id"`
	Topic   string `json:"topic"`
	Feeling string `json:"feeling"`
}

func NewTopicCreateResponse(topic *domain.Topic) *TopicCreateResponse {
	return &TopicCreateResponse{
		ID:      topic.ID,
		Topic:   topic.Topic,
		Feeling: topic.Feeling,
	}
}

type TopicUpdateResponse struct {
	ID      int    `json:"id"`
	Topic   string `json:"topic"`
	Feeling string `json:"feeling"`
}

func NewTopicUpdateResponse(topic *domain.Topic) *TopicUpdateResponse {
	return &TopicUpdateResponse{
		ID:      topic.ID,
		Topic:   topic.Topic,
		Feeling: topic.Feeling,
	}
}
