package service

import (
	"context"
	"samsamoohooh-go-api/internal/domain"
)

var _ domain.PostService = (*PostService)(nil)

type PostService struct {
	postRepository domain.PostRepository
}

func NewPostService(postRepository domain.PostRepository) *PostService {
	return &PostService{postRepository: postRepository}
}

func (p *PostService) Create(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	return p.postRepository.Create(ctx, post)
}

func (p *PostService) List(ctx context.Context, offset, limit int) ([]*domain.Post, error) {
	return p.postRepository.List(ctx, offset, limit)
}

func (p *PostService) GetByID(ctx context.Context, id int) (*domain.Post, error) {
	return p.postRepository.GetByID(ctx, id)
}

func (p *PostService) GetCommentsByID(ctx context.Context, id, offset, limit int) ([]*domain.Comment, error) {
	return p.postRepository.GetCommentsByID(ctx, id, offset, limit)
}

func (p *PostService) Update(ctx context.Context, id int, post *domain.Post) (*domain.Post, error) {
	return p.postRepository.Update(ctx, id, post)
}

func (p *PostService) Delete(ctx context.Context, id int) error {
	return p.postRepository.Delete(ctx, id)
}
