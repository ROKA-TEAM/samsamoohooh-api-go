package presenter

import (
	"samsamoohooh-go-api/internal/application/domain"
	"time"
)

type GroupCreateResponse struct {
	ID          int    `json:"id"`
	BookTitle   string `json:"bookTitle"`
	Author      string `json:"author"`
	MaxPage     int    `json:"maxPage"`
	Publisher   string `json:"publisher"`
	Description string `json:"description"`
	Bookmark    int    `json:"bookmark"`
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
	}
}

type GroupGetUsersByGroupIDResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Resolution string `json:"resolution"`
}

func NewGroupGetUsersByGroupIDResponse(users []*domain.User) []*GroupGetUsersByGroupIDResponse {
	var listUser []*GroupGetUsersByGroupIDResponse
	for _, user := range users {
		listUser = append(listUser, &GroupGetUsersByGroupIDResponse{
			ID:         user.ID,
			Name:       user.Name,
			Resolution: user.Resolution,
		})
	}

	return listUser
}

type GroupGetPostsByGroupIDResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewGroupGetPostsByGroupIDResponse(posts []*domain.Post) []*GroupGetPostsByGroupIDResponse {
	var listPost []*GroupGetPostsByGroupIDResponse
	for _, post := range posts {
		listPost = append(listPost, &GroupGetPostsByGroupIDResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
		})
	}

	return listPost
}

type GroupGetTasksByGroupIDResponse struct {
	ID        int       `json:"id"`
	Deadline  time.Time `json:"deadline"`
	Range     int       `json:"range"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewGroupGetTasksByIDResponse(tasks []*domain.Task) []*GroupGetTasksByGroupIDResponse {
	var listTask []*GroupGetTasksByGroupIDResponse
	for _, task := range tasks {
		listTask = append(listTask, &GroupGetTasksByGroupIDResponse{
			ID:        task.ID,
			Deadline:  task.Deadline,
			Range:     task.Range,
			CreatedAt: task.CreatedAt,
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

type GroupGenerateJoinCodeResponse struct {
	Code string `json:"code"`
}

func NewGroupGenerateJoinCodeResponse(code string) *GroupGenerateJoinCodeResponse {
	return &GroupGenerateJoinCodeResponse{
		Code: code,
	}
}

type GroupStartDiscussionResponse struct {
	Topics []string `json:"topics"`
	Users  []string `json:"users"`
}

func NewGroupStartDiscussionResponse(topics []string, users []string) *GroupStartDiscussionResponse {
	return &GroupStartDiscussionResponse{
		Topics: topics,
		Users:  users,
	}
}
