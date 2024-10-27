package presenter

import (
	"samsamoohooh-go-api/internal/application/domain"
	"time"
)

type TaskCreateRequest struct {
	Deadline time.Time `json:"deadline"`
	Range    int       `json:"range"`
	GroupID  int       `json:"groupID"`
}

func (r TaskCreateRequest) ToDomain() *domain.Task {
	return &domain.Task{
		Deadline: r.Deadline,
		Range:    r.Range,
		GroupID:  r.GroupID,
	}
}

type TaskGetTopicsByTaskIDRequest struct {
	ID     int `uri:"id"`
	Offset int `query:"offset"`
	Limit  int `query:"limit"`
}

type TaskUpdateRequest struct {
	ID       int       `uri:"id"`
	Deadline time.Time `json:"deadline,omitempty" validate:"omitempty"`
	Range    int       `json:"range,omitempty" validate:"omitempty"`
}

func (r TaskUpdateRequest) ToDomain() *domain.Task {
	return &domain.Task{
		Deadline: r.Deadline,
		Range:    r.Range,
	}
}

type TaskDeleteRequest struct {
	ID int `uri:"id"`
}
