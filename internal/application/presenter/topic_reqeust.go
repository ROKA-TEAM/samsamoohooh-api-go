package presenter

import (
	"samsamoohooh-go-api/internal/application/domain"
)

type TopicCreateRequest struct {
	Topic   string `json:"topic" validate:"min=1,max=24"`
	Feeling string `json:"feeling" validate:"min=1,max=500"`
	TaskID  int    `json:"taskID"`
}

func (r TopicCreateRequest) ToDomain() *domain.Topic {
	return &domain.Topic{
		Topic:   r.Topic,
		Feeling: r.Feeling,
		TaskID:  r.TaskID,
	}
}

type TopicUpdateRequest struct {
	Topic   string `json:"topic" validate:"min=1,max=24"`
	Feeling string `json:"feeling" validate:"min=1,max=500"`
}

func (r TopicUpdateRequest) ToDomain() *domain.Topic {
	return &domain.Topic{
		Topic:   r.Topic,
		Feeling: r.Feeling,
	}
}
