package service

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/port"
)

var _ port.UserService = (*UserService)(nil)

type UserService struct {
	userRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	createdUser, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *UserService) GetByUserID(ctx context.Context, id int) (*domain.User, error) {
	gotUser, err := s.userRepository.GetByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	return gotUser, nil
}

func (s *UserService) GetByUserSub(ctx context.Context, sub string) (*domain.User, error) {
	gotUser, err := s.userRepository.GetByUserSub(ctx, sub)
	if err != nil {
		return nil, err
	}

	return gotUser, nil
}

func (s *UserService) GetGroupsByUserID(ctx context.Context, id int, limit, offset int) ([]*domain.Group, error) {
	gotGroups, err := s.userRepository.GetGroupsByUserID(ctx, id, limit, offset)
	if err != nil {
		return nil, err
	}

	return gotGroups, nil
}

func (s *UserService) GetUsers(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	gotUsers, err := s.userRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return gotUsers, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int, user *domain.User) (*domain.User, error) {
	updatedUser, err := s.userRepository.UpdateUser(ctx, id, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	err := s.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) IsUserInGroup(ctx context.Context, userID, groupID int) (bool, error) {
	isIn, err := s.userRepository.IsUserInGroup(ctx, userID, groupID)
	if err != nil {
		return false, err
	}

	return isIn, nil
}
