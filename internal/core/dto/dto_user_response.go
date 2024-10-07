package dto

import (
	"samsamoohooh-go-api/internal/core/domain"
	"time"
)

type UserCreateResponse struct {
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

func NewUserCreateResponse(user *domain.User) *UserCreateResponse {
	return &UserCreateResponse{
		ID:         user.ID,
		Name:       user.Name,
		Resolution: user.Resolution,
		Role:       string(user.Role),
		Sub:        user.Sub,
		Social:     string(user.Social),
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		DeletedAt:  user.DeletedAt.Time,
	}
}

type UserGetByIDResponse struct {
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

func NewUserGetByIDResponse(user *domain.User) *UserGetByIDResponse {
	return &UserGetByIDResponse{
		ID:         user.ID,
		Name:       user.Name,
		Resolution: user.Resolution,
		Role:       string(user.Role),
		Sub:        user.Sub,
		Social:     string(user.Social),
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		DeletedAt:  user.DeletedAt.Time,
	}
}

type UserGetGroupsByIDResponse struct {
	ID          uint   `json:"id"`
	BookTitle   string `json:"bookTitle"`
	MaxPage     int    `json:"maxPage"`
	Publisher   string `json:"publisher"`
	Description string `json:"description"`
	Bookmark    int    `json:"bookMark"`
}

func NewUserGetGroupsByIDResponse(domainGroups []*domain.Group) []*UserGetGroupsByIDResponse {
	groups := make([]*UserGetGroupsByIDResponse, len(domainGroups))
	for _, g := range domainGroups {
		groups = append(groups, &UserGetGroupsByIDResponse{
			ID:          g.ID,
			BookTitle:   g.BookTitle,
			MaxPage:     g.MaxPage,
			Publisher:   g.Publisher,
			Description: g.Description,
			Bookmark:    g.Bookmark,
		})
	}

	return groups
}

type UserGetAllResponse struct {
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

func NewUserGetAllResponse(domainUsers []domain.User) []*UserGetAllResponse {
	users := make([]*UserGetAllResponse, len(domainUsers))
	for _, user := range domainUsers {
		users = append(users, &UserGetAllResponse{
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

type UserUpdateResponse struct {
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

func NewUserUpdateResponse(user *domain.User) *UserUpdateResponse {
	return &UserUpdateResponse{
		ID:         user.ID,
		Name:       user.Name,
		Resolution: user.Resolution,
		Role:       string(user.Role),
		Sub:        user.Sub,
		Social:     string(user.Social),
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		DeletedAt:  user.DeletedAt.Time,
	}
}

type UserDeleteResponse struct {
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

func NewUserDeleteResponse(user *domain.User) *UserDeleteResponse {
	return &UserDeleteResponse{
		ID:         user.ID,
		Name:       user.Name,
		Resolution: user.Resolution,
		Role:       string(user.Role),
		Sub:        user.Sub,
		Social:     string(user.Social),
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		DeletedAt:  user.DeletedAt.Time,
	}
}
