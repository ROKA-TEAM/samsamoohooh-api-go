package presenter

import (
	"samsamoohooh-go-api/internal/application/domain"
)

type GroupCreateReqeust struct {
	BookTitle   string `json:"bookTitle" validate:"min=1,max=28"`
	Author      string `json:"author" validate:"min=1,max=12,omitempty"`
	MaxPage     int    `json:"maxPage" validate:"min=1,max=9999,omitempty"`
	Publisher   string `json:"publisher" validate:"min=1,max=30,omitempty"`
	Description string `json:"description" validate:"min=0,max=56,omitempty"`
}

func (r *GroupCreateReqeust) ToDomain() *domain.Group {
	return &domain.Group{
		BookTitle:   r.BookTitle,
		Author:      r.Author,
		MaxPage:     r.MaxPage,
		Publisher:   r.Publisher,
		Description: r.Description,
	}
}

type GroupUpdateRequest struct {
	BookTitle   string `json:"bookTitle" validate:"min=1,max=28,omitempty"`
	Author      string `json:"author" validate:"min=1,max=12,omitempty"`
	MaxPage     int    `json:"maxPage" validate:"min=1,max=9999,omitempty"`
	Publisher   string `json:"publisher" validate:"min=1,max=30,omitempty"`
	Description string `json:"description" validate:"min=0,max=56,omitempty"`
}

func (r *GroupUpdateRequest) ToDomain() *domain.Group {
	return &domain.Group{
		BookTitle:   r.BookTitle,
		Author:      r.Author,
		MaxPage:     r.MaxPage,
		Publisher:   r.Publisher,
		Description: r.Description,
	}
}
