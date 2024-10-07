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

type GroupGetByIDResponse struct {
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

func NewGroupGetByIDResponse(g *domain.Group) *GroupGetByIDResponse {
	return &GroupGetByIDResponse{
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

type GroupGetUsersByIDResponse struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Resolution string    `json:"resolution"`
	Role       string    `json:"role"`
	Sub        string    `json:"sub"`
	Social     string    `json:"social"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at,omitempty"`
}

func NewGroupsGetUsersByIDResponse(domainUsers []*domain.User) []*GroupGetUsersByIDResponse {
	users := make([]*GroupGetUsersByIDResponse, len(domainUsers))
	for _, user := range domainUsers {
		users = append(users, &GroupGetUsersByIDResponse{
			ID:         user.ID,
			Name:       user.Name,
			Resolution: user.Resolution,
			Role:       string(user.Role),
			Sub:        user.Sub,
			Social:     string(user.Social),
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
			DeletedAt:  user.DeletedAt.Time,
		})
	}

	return users
}
