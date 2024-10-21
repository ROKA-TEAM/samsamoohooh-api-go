package service

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
)

var _ domain.CommentService = (*CommentService)(nil)

type CommentService struct {
	commentRepository domain.CommentRepository
}

func NewCommentService(
	commentRepository domain.CommentRepository,
) *CommentService {
	return &CommentService{commentRepository: commentRepository}
}

func (s *CommentService) CreateComment(ctx context.Context, postID int, comment *domain.Comment) (*domain.Comment, error) {
	createdComment, err := s.commentRepository.CreateComment(ctx, postID, comment)
	if err != nil {
		return nil, err
	}

	return createdComment, nil
}

func (s *CommentService) GetComments(ctx context.Context, offset, limit int) ([]*domain.Comment, error) {
	gotComment, err := s.commentRepository.GetComments(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	return gotComment, nil
}

func (s *CommentService) GetByCommentID(ctx context.Context, id int) (*domain.Comment, error) {
	gotComment, err := s.commentRepository.GetByCommentID(ctx, id)
	if err != nil {
		return nil, err
	}

	return gotComment, nil
}

func (s *CommentService) UpdateComment(ctx context.Context, id int, comment *domain.Comment) (*domain.Comment, error) {
	updatedComment, err := s.commentRepository.UpdateComment(ctx, id, comment)
	if err != nil {
		return nil, err
	}

	return updatedComment, nil
}

func (s *CommentService) DeleteComment(ctx context.Context, id int) error {
	err := s.commentRepository.DeleteComment(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
