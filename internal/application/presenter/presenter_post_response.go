package presenter

import (
	"samsamoohooh-go-api/internal/application/domain"
	"time"
)

type PostCreateResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewPostCreateResponse(post *domain.Post) *PostCreateResponse {
	return &PostCreateResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
	}
}

type PostGetCommentsByPostIDResponse struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewPostGetCommentsByIDResponse(comments []*domain.Comment) []*PostGetCommentsByPostIDResponse {
	var listComment []*PostGetCommentsByPostIDResponse
	for _, comment := range comments {
		listComment = append(listComment, &PostGetCommentsByPostIDResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
		})
	}

	return listComment
}

type PostUpdateResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewPostUpdateResponse(post *domain.Post) *PostUpdateResponse {
	return &PostUpdateResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}
