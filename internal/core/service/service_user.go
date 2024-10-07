package service

import (
	"context"
	"samsamoohooh-go-api/internal/core/dto"
	"samsamoohooh-go-api/internal/core/port"
)

var _ port.UserService = (*UserService)(nil)

type UserService struct {
	userRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) Create(ctx context.Context, user *dto.UserCreateRequest) (*dto.UserCreateResponse, error) {
	u, err := s.userRepository.Create(ctx, user.ToDomain())
	if err != nil {
		return nil, err
	}

	return dto.NewUserCreateResponse(u), nil
}

func (s *UserService) GetByID(ctx context.Context, id uint) (*dto.UserGetByIDResponse, error) {
	u, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.NewUserGetByIDResponse(u), nil
}

func (s *UserService) GetGroupsByID(ctx context.Context, id uint) ([]*dto.UserGetGroupsByIDResponse, error) {
	groups, err := s.userRepository.GetGroupsByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.NewUserGetGroupsByIDResponse(groups), nil
}

func (s *UserService) GetAll(ctx context.Context, skip, limit int) ([]*dto.UserGetAllResponse, error) {
	users, err := s.userRepository.GetAll(ctx, skip, limit)
	if err != nil {
		return nil, err
	}

	return dto.NewUserGetAllResponse(users), nil

}

func (s *UserService) Update(ctx context.Context, id uint, user *dto.UserUpdateRequest) (*dto.UserUpdateResponse, error) {
	update, err := s.userRepository.Update(ctx, id, user.ToDomain())
	if err != nil {
		return nil, err
	}

	return dto.NewUserUpdateResponse(update), nil
}

func (s *UserService) Delete(ctx context.Context, id uint) (*dto.UserDeleteResponse, error) {
	user, err := s.userRepository.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.NewUserDeleteResponse(user), nil
}
