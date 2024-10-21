package service

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
)

type CommentService struct {
	commentRepository domain.CommentRepository
}

func NewCommentService(commentRepository domain.CommentRepository) *CommentService {
	return &CommentService{commentRepository: commentRepository}
}

func (s *CommentService) Create(ctx context.Context, postID int, comment *domain.Comment) (*domain.Comment, error) {
	return s.commentRepository.Create(ctx, postID, comment)
}

func (s *CommentService) List(ctx context.Context, offset, limit int) ([]*domain.Comment, error) {
	return s.commentRepository.List(ctx, offset, limit)
}

func (s *CommentService) GetByID(ctx context.Context, id int) (*domain.Comment, error) {
	return s.commentRepository.GetByID(ctx, id)
}

func (s *CommentService) Update(ctx context.Context, id int, comment *domain.Comment) (*domain.Comment, error) {
	return s.commentRepository.Update(ctx, id, comment)
}

func (s *CommentService) Delete(ctx context.Context, id int) error {
	return s.commentRepository.Delete(ctx, id)
}
