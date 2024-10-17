package presenter

import "samsamoohooh-go-api/internal/domain"

type PostCreateRequest struct {
	Title   string `json:"title" validate:"min=1,max=24"`
	Content string `json:"content" validate:"min=1,max=300"`
}

func (r PostCreateRequest) ToDomain() *domain.Post {
	return &domain.Post{
		Title:   r.Title,
		Content: r.Content,
	}
}

type PostUpdateRequest struct {
	Title   string `json:"title" validate:"min=1,max=24,omitempty"`
	Content string `json:"content" validate:"min=1,max=300,omitempty"`
}

func (r PostUpdateRequest) ToDomain() *domain.Post {
	return &domain.Post{
		Title:   r.Title,
		Content: r.Content,
	}
}
