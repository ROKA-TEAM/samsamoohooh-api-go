package presenter

import (
	"samsamoohooh-go-api/internal/domain"
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

type TaskUdpateRequest struct {
	Deadline time.Time `json:"deadline,omitempty" validate:"omitempty"`
	Range    int       `json:"range,omitempty" validate:"omitempty"`
}

func (r TaskUdpateRequest) ToDomain() *domain.Task {
	return &domain.Task{
		Deadline: r.Deadline,
		Range:    r.Range,
	}
}
