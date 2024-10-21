package service

import (
	"context"
	domain "samsamoohooh-go-api/internal/application/domain"
)

var _ domain.PostService = (*PostService)(nil)

type PostService struct {
	postRepository domain.PostRepository
}

func NewPostService(
	postRepository domain.PostRepository,
) *PostService {
	return &PostService{postRepository: postRepository}
}

func (p *PostService) CreatePost(ctx context.Context, groupID int, post *domain.Post) (*domain.Post, error) {
	createdPost, err := p.postRepository.CreatePost(ctx, groupID, post)
	if err != nil {
		return nil, err
	}

	return createdPost, nil
}

func (p *PostService) GetPosts(ctx context.Context, offset, limit int) ([]*domain.Post, error) {
	listPost, err := p.postRepository.GetPosts(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	return listPost, nil
}

func (p *PostService) GetByPostID(ctx context.Context, id int) (*domain.Post, error) {
	gotPost, err := p.postRepository.GetByPostID(ctx, id)
	if err != nil {
		return nil, err
	}

	return gotPost, nil
}

func (p *PostService) GetCommentsByPostID(ctx context.Context, id, offset, limit int) ([]*domain.Comment, error) {
	listComment, err := p.postRepository.GetCommentsByPostID(ctx, id, offset, limit)
	if err != nil {
		return nil, err
	}

	return listComment, nil
}

func (p *PostService) UpdatePost(ctx context.Context, id int, post *domain.Post) (*domain.Post, error) {
	updatedPost, err := p.postRepository.UpdatePost(ctx, id, post)
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}

func (p *PostService) DeletePost(ctx context.Context, id int) error {
	err := p.postRepository.DeletePost(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
