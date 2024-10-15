package service

import (
	"context"
	"samsamoohooh-go-api/internal/domain"
)

type UserService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	createdUser, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *UserService) GetByID(ctx context.Context, id int) (*domain.User, error) {
	gotUser, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return gotUser, nil
}

func (s *UserService) GetGroupsByID(ctx context.Context, id int) ([]*domain.Group, error) {
	gotGroups, err := s.userRepository.GetGroupsByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return gotGroups, nil
}

func (s *UserService) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	gotUsers, err := s.userRepository.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return gotUsers, nil
}

func (s *UserService) Update(ctx context.Context, id int, user *domain.User) (*domain.User, error) {
	updatedUser, err := s.userRepository.Update(ctx, id, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	err := s.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
