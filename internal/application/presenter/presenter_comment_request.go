package presenter

import "samsamoohooh-go-api/internal/application/domain"

type CommentCreateRequest struct {
	PostID  int    `json:"postID"`
	Content string `json:"content" validate:"min=1,max=24"`
}

func (r CommentCreateRequest) ToDomain() *domain.Comment {
	return &domain.Comment{
		Content: r.Content,
		PostID:  r.PostID,
	}
}

type CommentUpdateRequest struct {
	ID      int    `uri:"id"`
	Content string `json:"content" validate:"min=1,max=24"`
}

func (r CommentUpdateRequest) ToDomain() *domain.Comment {
	return &domain.Comment{
		Content: r.Content,
	}
}

type CommentDeleteRequest struct {
	ID int `uri:"id"`
}
