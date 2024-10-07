package dto

import "samsamoohooh-go-api/internal/core/domain"

type GroupCreateRequest struct {
	BookTitle   string `json:"bookTitle"`
	Author      string `json:"author"`
	MaxPage     int    `json:"maxPage"`
	Publisher   string `json:"publisher"`
	Description string `json:"description"`
	Bookmark    int    `json:"bookmark"`
}

func (r *GroupCreateRequest) ToDomain() *domain.Group {
	return &domain.Group{
		BookTitle:   r.BookTitle,
		Author:      r.Author,
		MaxPage:     r.MaxPage,
		Publisher:   r.Publisher,
		Description: r.Description,
		Bookmark:    r.Bookmark,
	}
}
