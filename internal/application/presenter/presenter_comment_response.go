package presenter

import (
	"samsamoohooh-go-api/internal/application/domain"
	"time"
)

type CommentCreateResponse struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewCommentCreateResponse(comment *domain.Comment) *CommentCreateResponse {
	return &CommentCreateResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
	}
}

type CommentUpdateResponse struct {
	Content string `json:"content"`
}

func NewCommentUpdateResponse(comment *domain.Comment) *CommentUpdateResponse {
	return &CommentUpdateResponse{
		Content: comment.Content,
	}
}
