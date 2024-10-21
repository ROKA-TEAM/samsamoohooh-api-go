package service

import (
	"context"
	domain2 "samsamoohooh-go-api/internal/application/domain"
)

var _ domain2.PostService = (*PostService)(nil)

type PostService struct {
	postRepository domain2.PostRepository
}

func NewPostService(postRepository domain2.PostRepository) *PostService {
	return &PostService{postRepository: postRepository}
}

func (p *PostService) Create(ctx context.Context, groupID int, post *domain2.Post) (*domain2.Post, error) {
	return p.postRepository.Create(ctx, groupID, post)
}

func (p *PostService) List(ctx context.Context, offset, limit int) ([]*domain2.Post, error) {
	return p.postRepository.List(ctx, offset, limit)
}

func (p *PostService) GetByID(ctx context.Context, id int) (*domain2.Post, error) {
	return p.postRepository.GetByID(ctx, id)
}

func (p *PostService) GetCommentsByID(ctx context.Context, id, offset, limit int) ([]*domain2.Comment, error) {
	return p.postRepository.GetCommentsByID(ctx, id, offset, limit)
}

func (p *PostService) Update(ctx context.Context, id int, post *domain2.Post) (*domain2.Post, error) {
	return p.postRepository.Update(ctx, id, post)
}

func (p *PostService) Delete(ctx context.Context, id int) error {
	return p.postRepository.Delete(ctx, id)
}
