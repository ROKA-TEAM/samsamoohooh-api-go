package service

import (
	"context"
	"samsamoohooh-go-api/internal/domain"

	"github.com/pkg/errors"
)

var _ domain.GroupService = &GroupService{}

type GroupService struct {
	groupRepository domain.GroupRepository
	userService     domain.UserService
	taskService     domain.TaskService
}

func NewGroupService(groupRepository domain.GroupRepository) *GroupService {
	return &GroupService{
		groupRepository: groupRepository,
	}
}

func (s *GroupService) Create(ctx context.Context, group *domain.Group) (*domain.Group, error) {
	token, ok := ctx.Value("token").(*domain.Token)
	if !ok {
		return nil, errors.Wrap(domain.ErrInternal, "token value cannot be converted")
	}
	return s.groupRepository.Create(ctx, token.Subject, group)
}
func (s *GroupService) List(ctx context.Context, offset, limit int) ([]*domain.Group, error) {
	return s.groupRepository.List(ctx, offset, limit)
}
func (s *GroupService) GetByID(ctx context.Context, id int) (*domain.Group, error) {
	return s.groupRepository.GetByID(ctx, id)
}

func (s *GroupService) GetUsersByID(ctx context.Context, id int, offset, limit int) ([]*domain.User, error) {
	return s.groupRepository.GetUsersByID(ctx, id, offset, limit)
}
func (s *GroupService) GetPostsByID(ctx context.Context, id int, offset, limit int) ([]*domain.Post, error) {
	return s.groupRepository.GetPostsByID(ctx, id, offset, limit)
}
func (s *GroupService) GetTasksByID(ctx context.Context, id int, offset, limit int) ([]*domain.Task, error) {
	return s.groupRepository.GetTasksByID(ctx, id, offset, limit)
}
func (s *GroupService) Update(ctx context.Context, id int, group *domain.Group) (*domain.Group, error) {
	return s.groupRepository.Update(ctx, id, group)
}
func (s *GroupService) Delete(ctx context.Context, id int) error {
	return s.groupRepository.Delete(ctx, id)
}

func (s *GroupService) StartDiscussion(ctx context.Context, groupID, taskID int) (topics []string, userNames []string, err error) {
	token, ok := ctx.Value("token").(*domain.Token)
	if !ok {
		return nil, nil, errors.Wrap(domain.ErrInternal, "token value cannot be converted")
	}

	// 요청한 사용자가 조회해도 되나?
	_, err = s.userService.GetGroupsByID(ctx, token.Subject, 0, 10)
	if err != nil {
		return nil, nil, errors.Wrap(domain.ErrConstraint, "user cannot access the group")
	}

	// topics 구하기
	// topic == group'users.len
	usersLen, err := s.groupRepository.GetUsersLenByID(ctx, groupID)
	if err != nil {
		return nil, nil, err
	}

	queriedTopics, err := s.taskService.GetTopicsByID(ctx, taskID, 0, usersLen)
	if err != nil {
		return nil, nil, err
	}

	// users 구하기
	queriedUsers, err := s.groupRepository.GetUsersByID(ctx, groupID, 0, usersLen)
	if err != nil {
		return nil, nil, err
	}

	for _, queryTopic := range queriedTopics {
		topics = append(topics, queryTopic.Topic)
	}

	for _, queryUser := range queriedUsers {
		userNames = append(userNames, queryUser.Name)
	}

	return topics, userNames, nil
}
