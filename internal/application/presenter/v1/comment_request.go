package v1

import (
	"samsamoohooh-go-api/internal/application/domain"
)

type CommentCreateRequest struct {
	Content string `json:"content" validate:"min=1,max=24"`
	PostID  int    `json:"postID"`
}

func (r CommentCreateRequest) ToDomain() *domain.Comment {
	return &domain.Comment{
		Content: r.Content,
		PostID:  r.PostID,
	}
}

type CommentUpdateRequest struct {
	Content string `json:"content" validate:"min=1,max=24"`
}

func (r CommentUpdateRequest) ToDomain() *domain.Comment {
	return &domain.Comment{
		Content: r.Content,
	}
}
