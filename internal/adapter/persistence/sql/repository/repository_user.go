package repository

import (
	"context"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/database"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/repository/utils"
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
		return nil, utils.Wrap(err)
	}

	return user, nil
}

func (r *UserRepository) GetByUserID(ctx context.Context, id uint) (*domain.User, error) {
	user := domain.User{}
	err := r.database.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return &user, nil
}

func (r *UserRepository) GetGroupsByUserID(ctx context.Context, id int) ([]*domain.Group, error) {
	user := domain.User{}
	err := r.database.WithContext(ctx).Preload("Groups").First(&user, id).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return user.Groups, nil
}

func (r *UserRepository) GetAll(ctx context.Context, skip, limit int) ([]*domain.User, error) {
	users := []*domain.User{}
	err := r.database.Limit(limit).Offset((skip - 1) * limit).Find(&users).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, id uint, user *domain.User) (*domain.User, error) {
	user.Model.ID = id
	err := r.database.Save(&user).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	err := r.database.Delete(&domain.User{}, id).Error
	if err != nil {
		return utils.Wrap(err)
	}

	return nil
}
