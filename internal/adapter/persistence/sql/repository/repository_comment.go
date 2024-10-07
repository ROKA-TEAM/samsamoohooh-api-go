package repository

import (
	"context"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/database"
	"samsamoohooh-go-api/internal/core/domain"
	"samsamoohooh-go-api/internal/core/port"
)

var _ port.CommentRepository = (*CommentRepository)(nil)

type CommentRepository struct {
	database *database.Database
}

func NewCommentRepository(database *database.Database) *CommentRepository {
	return &CommentRepository{
		database: database,
	}
}

func (r *CommentRepository) Create(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	err := r.database.WithContext(ctx).Create(comment).Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *CommentRepository) GetByID(ctx context.Context, id uint) (*domain.Comment, error) {
	comment := domain.Comment{}
	err := r.database.WithContext(ctx).First(&comment, id).Error
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *CommentRepository) GetAll(ctx context.Context, skip, limit int) ([]domain.Comment, error) {
	var comments []domain.Comment
	err := r.database.WithContext(ctx).Limit(limit).Offset((skip - 1) * limit).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil

}

func (r *CommentRepository) Update(ctx context.Context, id uint, comment *domain.Comment) (*domain.Comment, error) {
	comment.ID = id
	err := r.database.WithContext(ctx).Save(comment).Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *CommentRepository) Delete(ctx context.Context, id uint) error {
	err := r.database.WithContext(ctx).Delete(&domain.Comment{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
