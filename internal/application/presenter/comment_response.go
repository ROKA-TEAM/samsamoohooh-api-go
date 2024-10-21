package presenter

import (
	"samsamoohooh-go-api/internal/application/domain"
	"time"
)

type CommentCreateResponse struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCommentCreateResponse(comment *domain.Comment) *CommentCreateResponse {
	return &CommentCreateResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}

type CommentListResponse struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCommentListResponse(comments []*domain.Comment) []*CommentListResponse {
	var listComments []*CommentListResponse
	for _, comment := range comments {
		listComments = append(listComments, &CommentListResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}

	return listComments
}

type CommentGetByCommentID struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCommentGetByCommentIDResponse(comment *domain.Comment) *CommentGetByCommentID {
	return &CommentGetByCommentID{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}

type CommentUpdateResponse struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCommentUpdateResponse(comment *domain.Comment) *CommentUpdateResponse {
	return &CommentUpdateResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}
