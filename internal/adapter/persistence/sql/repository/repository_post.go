package repository

import (
	"context"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/database"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/repository/utils"
	"samsamoohooh-go-api/internal/core/domain"
	"samsamoohooh-go-api/internal/core/port"
)

var _ port.PostRepository = (*PostRepository)(nil)

type PostRepository struct {
	database *database.Database
}

func NewPostRepository(database *database.Database) *PostRepository {
	return &PostRepository{
		database: database,
	}
}

func (r *PostRepository) Create(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	err := r.database.WithContext(ctx).Create(post).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return post, nil
}

func (r *PostRepository) GetByID(ctx context.Context, id uint) (*domain.Post, error) {
	post := domain.Post{}
	err := r.database.WithContext(ctx).First(&post, id).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return &post, nil
}

func (r *PostRepository) GetCommentsByID(ctx context.Context, id uint) ([]*domain.Comment, error) {
	comments := []*domain.Comment{}
	err := r.database.WithContext(ctx).Preload("Comments").First(comments, id).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return comments, nil
}

func (r *PostRepository) GetAll(ctx context.Context, skip, limit int) ([]*domain.Post, error) {
	posts := []*domain.Post{}
	err := r.database.WithContext(ctx).Limit(limit).Offset((skip - 1) * limit).Find(&posts).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return posts, nil
}

func (r *PostRepository) Update(ctx context.Context, id uint, user *domain.Post) (*domain.Post, error) {
	user.Model.ID = id
	err := r.database.WithContext(ctx).Save(user).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return user, nil
}

func (r *PostRepository) Delete(ctx context.Context, id uint) error {
	err := r.database.WithContext(ctx).Delete(&domain.Post{}, id).Error
	if err != nil {
		return utils.Wrap(err)
	}

	return nil
}
