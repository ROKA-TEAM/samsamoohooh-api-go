package repository

import (
	"context"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/database"
	"samsamoohooh-go-api/internal/core/domain"
	"samsamoohooh-go-api/internal/core/port"
)

var _ port.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	database *database.Database
}

func NewUserRepository(database *database.Database) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := r.database.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	user := domain.User{}
	err := r.database.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetGroupsByID(ctx context.Context, id uint) ([]*domain.Group, error) {
	user := domain.User{}
	err := r.database.WithContext(ctx).Preload("Groups").First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return user.Groups, nil
}

func (r *UserRepository) GetAll(ctx context.Context, skip, limit int) ([]domain.User, error) {
	var users []domain.User
	err := r.database.WithContext(ctx).Limit(limit).Offset((skip - 1) * limit).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, id uint, user *domain.User) (*domain.User, error) {
	user.ID = id
	err := r.database.WithContext(ctx).Save(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint) (*domain.User, error) {
	err := r.database.WithContext(ctx).Delete(&domain.User{}, id).Error
	if err != nil {
		return nil, err
	}

	deletedUser := &domain.User{}
	if err := r.database.WithContext(ctx).Unscoped().Where("id = ? AND deleted_at IS NOT NULL", id).First(deletedUser).Error; err != nil {
		return nil, err
	}

	return deletedUser, nil
}
