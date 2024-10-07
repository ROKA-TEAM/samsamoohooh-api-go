package dto

import (
	"samsamoohooh-go-api/internal/core/domain"
	"time"
)

type GroupCreateResponse struct {
	ID          uint      `json:"id"`
	BookTitle   string    `json:"bookTitle"`
	Author      string    `json:"author"`
	MaxPage     int       `json:"maxPage"`
	Publisher   string    `json:"publisher"`
	Description string    `json:"description"`
	Bookmark    int       `json:"bookmark"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

func NewGroupCreateResponse(g *domain.Group) *GroupCreateResponse {
	return &GroupCreateResponse{
		ID:          g.ID,
		BookTitle:   g.BookTitle,
		Author:      g.Author,
		MaxPage:     g.MaxPage,
		Publisher:   g.Publisher,
		Description: g.Description,
		Bookmark:    g.Bookmark,
		CreatedAt:   g.CreatedAt,
		UpdatedAt:   g.UpdatedAt,
		DeletedAt:   g.DeletedAt.Time,
	}
}
