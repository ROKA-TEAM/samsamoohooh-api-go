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

func (r *UserRepository) Create(user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (r *UserRepository) GetByUserID(ctx context.Context, id uint) (*domain.User, error) {
	return nil, nil
}
func (r *UserRepository) GetGroupsByUserID(ctx context.Context, id int) ([]*domain.Group, error) {
	return nil, nil
}
func (r *UserRepository) GetAll(ctx context.Context, skip, limit int) ([]*domain.User, error) {
	return nil, nil
}

func (r *UserRepository) Update(ctx context.Context, id uint, user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	return nil
}
