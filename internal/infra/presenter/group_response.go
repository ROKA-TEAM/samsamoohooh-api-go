package presenter

import (
	"samsamoohooh-go-api/internal/domain"
	"time"
)

type GroupCreateResponse struct {
	ID          int       `json:"id"`
	BookTitle   string    `json:"bookTitle"`
	Author      string    `json:"author"`
	MaxPage     int       `json:"maxPage"`
	Publisher   string    `json:"publisher"`
	Description string    `json:"description"`
	Bookmark    int       `json:"bookmark"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewGroupCreateResponse(group *domain.Group) *GroupCreateResponse {
	return &GroupCreateResponse{
		ID:          group.ID,
		BookTitle:   group.BookTitle,
		Author:      group.Author,
		MaxPage:     group.MaxPage,
		Publisher:   group.Publisher,
		Description: group.Description,
		Bookmark:    group.Bookmark,
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
	}
}

type GroupListResponse struct {
	ID          int       `json:"id"`
	BookTitle   string    `json:"bookTitle"`
	Author      string    `json:"author"`
	MaxPage     int       `json:"maxPage"`
	Publisher   string    `json:"publisher"`
	Description string    `json:"description"`
	Bookmark    int       `json:"bookmark"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewGroupListResponse(groups []*domain.Group) []*GroupListResponse {
	var response []*GroupListResponse
	for _, group := range groups {
		response = append(response, &GroupListResponse{
			ID:          group.ID,
			BookTitle:   group.BookTitle,
			Author:      group.Author,
			MaxPage:     group.MaxPage,
			Publisher:   group.Publisher,
			Description: group.Description,
			Bookmark:    group.Bookmark,
			CreatedAt:   group.CreatedAt,
			UpdatedAt:   group.UpdatedAt,
		})
	}

	return response
}

type GroupGetByIDResponse struct {
	ID          int       `json:"id"`
	BookTitle   string    `json:"bookTitle"`
	Author      string    `json:"author"`
	MaxPage     int       `json:"maxPage"`
	Publisher   string    `json:"publisher"`
	Description string    `json:"description"`
	Bookmark    int       `json:"bookmark"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewGroupGetByIDResponse(group *domain.Group) *GroupGetByIDResponse {
	return &GroupGetByIDResponse{
		ID:          group.ID,
		BookTitle:   group.BookTitle,
		Author:      group.Author,
		MaxPage:     group.MaxPage,
		Publisher:   group.Publisher,
		Description: group.Description,
		Bookmark:    group.Bookmark,
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
	}
}

type GroupGetUsersByIDResponse struct {
	ID         int                   `json:"id"`
	Name       string                `json:"name"`
	Resolution string                `json:"resolution"`
	Role       domain.UserRoleType   `json:"role"`
	Social     domain.UserSocialType `json:"social"`
	SocialSub  string                `json:"socialSub"`
	CreatedAt  time.Time             `json:"createdAt"`
	UpdatedAt  time.Time             `json:"updatedAt"`
}

func NewGroupGetUsersByIDResponse(users []*domain.User) []*GroupGetUsersByIDResponse {
	var listUser []*GroupGetUsersByIDResponse
	for _, user := range users {
		listUser = append(listUser, &GroupGetUsersByIDResponse{
			ID:         user.ID,
			Name:       user.Name,
			Resolution: user.Resolution,
			Role:       user.Role,
			Social:     user.Social,
			SocialSub:  user.SocialSub,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		})
	}

	return listUser
}

type GroupGetPostsByIDResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewGroupGetPostsByIDResponse(posts []*domain.Post) []*GroupGetPostsByIDResponse {
	var listPost []*GroupGetPostsByIDResponse
	for _, post := range posts {
		listPost = append(listPost, &GroupGetPostsByIDResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return listPost
}

type GroupGetTasksByIDResponse struct {
	ID        int       `json:"id"`
	Deadline  time.Time `json:"deadline"`
	Range     int       `json:"range"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewGroupGetTasksByIDResponse(tasks []*domain.Task) []*GroupGetTasksByIDResponse {
	var listTask []*GroupGetTasksByIDResponse
	for _, task := range tasks {
		listTask = append(listTask, &GroupGetTasksByIDResponse{
			ID:        task.ID,
			Deadline:  task.Deadline,
			Range:     task.Range,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		})
	}

	return listTask
}

type GroupUpdateResponse struct {
	ID          int       `json:"id"`
	BookTitle   string    `json:"bookTitle"`
	Author      string    `json:"author"`
	MaxPage     int       `json:"maxPage"`
	Publisher   string    `json:"publisher"`
	Description string    `json:"description"`
	Bookmark    int       `json:"bookmark"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewGroupUpdateResponse(group *domain.Group) *GroupUpdateResponse {
	return &GroupUpdateResponse{
		ID:          group.ID,
		BookTitle:   group.BookTitle,
		Author:      group.Author,
		MaxPage:     group.MaxPage,
		Publisher:   group.Publisher,
		Description: group.Description,
		Bookmark:    group.Bookmark,
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
	}
}
