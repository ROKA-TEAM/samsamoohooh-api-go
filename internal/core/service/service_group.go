package service

import (
	"context"
	"samsamoohooh-go-api/internal/core/dto"
	"samsamoohooh-go-api/internal/core/port"
)

var _ port.UserService = (*UserService)(nil)

type GroupService struct {
	userRepository  port.UserRepository
	groupRepository port.GroupRepository
}

func NewGroupService(groupRepository port.GroupRepository, userRepository port.UserRepository) *GroupService {
	return &GroupService{userRepository: userRepository, groupRepository: groupRepository}
}

func (s *GroupService) Create(ctx context.Context, group *dto.GroupCreateRequest, creatorID uint) (*dto.GroupCreateResponse, error) {

	// 모임을 생성하고자 하는 사용자 쿼리
	creator, err := s.userRepository.GetByID(ctx, creatorID)
	if err != nil {
		return nil, err
	}

	// dto -> group domain
	domainGroup := group.ToDomain()

	// 사용자를 생성할 group 에 추가
	domainGroup.Users = append(domainGroup.Users, creator)

	createdGroup, err := s.groupRepository.Create(ctx, domainGroup)
	if err != nil {
		return nil, err
	}

	return dto.NewGroupCreateResponse(createdGroup), nil
}

func (s *GroupService) GetByID(ctx context.Context, id uint) (*dto.GroupGetByIDResponse, error) {
	queriedGroup, err := s.groupRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.NewGroupGetByIDResponse(queriedGroup), nil
}

func (s *GroupService) GetUsersByID(ctx context.Context, id uint) ([]*dto.GroupGetUsersByIDResponse, error) {

	queriedUsers, err := s.groupRepository.GetUsersByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.NewGroupsGetUsersByIDResponse(queriedUsers), nil
}
