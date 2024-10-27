package presenter

import (
	"samsamoohooh-go-api/internal/application/domain"
	"time"
)

type TaskCreateResponse struct {
	ID        int       `json:"id"`
	Deadline  time.Time `json:"deadline"`
	Range     int       `json:"range"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewTaskCreateResponse(task *domain.Task) *TaskCreateResponse {
	return &TaskCreateResponse{
		ID:       task.ID,
		Deadline: task.Deadline,
		Range:    task.Range,
	}
}

type TaskGetTopicsByTaskIDResponse struct {
	ID      int    `json:"id"`
	Topic   string `json:"topic"`
	Feeling string `json:"feeling"`
}

func NewTaskGetTopicsByIDResponse(topics []*domain.Topic) []*TaskGetTopicsByTaskIDResponse {
	var topicList []*TaskGetTopicsByTaskIDResponse
	for _, topic := range topics {
		topicList = append(topicList, &TaskGetTopicsByTaskIDResponse{
			ID:      topic.ID,
			Topic:   topic.Topic,
			Feeling: topic.Feeling,
		})
	}
	return topicList
}

type TaskUpdateResponse struct {
	ID       int       `json:"id"`
	Deadline time.Time `json:"deadline"`
	Range    int       `json:"range"`
}

func NewTaskUpdateResponse(task *domain.Task) *TaskUpdateResponse {
	return &TaskUpdateResponse{
		ID:       task.ID,
		Deadline: task.Deadline,
		Range:    task.Range,
	}
}
