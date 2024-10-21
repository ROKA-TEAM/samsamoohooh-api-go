package presenter

import (
	"samsamoohooh-go-api/internal/application/domain"
	"time"
)

type TopicCreateResponse struct {
	ID        int       `json:"id"`
	Topic     string    `json:"topic"`
	Feeling   string    `json:"feeling"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewTopicCreateResponse(topic *domain.Topic) *TopicCreateResponse {
	return &TopicCreateResponse{
		ID:        topic.ID,
		Topic:     topic.Topic,
		Feeling:   topic.Feeling,
		CreatedAt: topic.CreatedAt,
		UpdatedAt: topic.UpdatedAt,
	}
}

type TopicListResponse struct {
	ID        int       `json:"id"`
	Topic     string    `json:"topic"`
	Feeling   string    `json:"feeling"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewTopicListResponse(topics []*domain.Topic) []*TopicListResponse {
	var listTopicResp []*TopicListResponse
	for _, topic := range topics {
		listTopicResp = append(listTopicResp, &TopicListResponse{
			ID:        topic.ID,
			Topic:     topic.Topic,
			Feeling:   topic.Feeling,
			CreatedAt: topic.CreatedAt,
			UpdatedAt: topic.UpdatedAt,
		})
	}

	return listTopicResp
}

type TopicGetByIDResponse struct {
	ID        int       `json:"id"`
	Topic     string    `json:"topic"`
	Feeling   string    `json:"feeling"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewTopicGetByIDResponse(topic *domain.Topic) *TopicGetByIDResponse {
	return &TopicGetByIDResponse{
		ID:        topic.ID,
		Topic:     topic.Topic,
		Feeling:   topic.Feeling,
		CreatedAt: topic.CreatedAt,
		UpdatedAt: topic.UpdatedAt,
	}
}

type TopicUpdateResponse struct {
	ID        int       `json:"id"`
	Topic     string    `json:"topic"`
	Feeling   string    `json:"feeling"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewTopicUpdateResponse(topic *domain.Topic) *TopicUpdateResponse {
	return &TopicUpdateResponse{
		ID:        topic.ID,
		Topic:     topic.Topic,
		Feeling:   topic.Feeling,
		CreatedAt: topic.CreatedAt,
		UpdatedAt: topic.UpdatedAt,
	}
}
