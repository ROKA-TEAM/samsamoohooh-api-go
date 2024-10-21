package service

import (
	"context"
	"github.com/pkg/errors"
	domain2 "samsamoohooh-go-api/internal/application/domain"
)

var _ domain2.GroupService = &GroupService{}

type GroupService struct {
	groupRepository domain2.GroupRepository
	userService     domain2.UserService
	taskService     domain2.TaskService
}

func NewGroupService(groupRepository domain2.GroupRepository) *GroupService {
	return &GroupService{
		groupRepository: groupRepository,
	}
}

func (s *GroupService) Create(ctx context.Context, group *domain2.Group) (*domain2.Group, error) {
	token, ok := ctx.Value("token").(*domain2.Token)
	if !ok {
		return nil, errors.Wrap(domain2.ErrInternal, "token value cannot be converted")
	}
	return s.groupRepository.Create(ctx, token.Subject, group)
}
func (s *GroupService) List(ctx context.Context, offset, limit int) ([]*domain2.Group, error) {
	return s.groupRepository.List(ctx, offset, limit)
}
func (s *GroupService) GetByID(ctx context.Context, id int) (*domain2.Group, error) {
	return s.groupRepository.GetByID(ctx, id)
}

func (s *GroupService) GetUsersByID(ctx context.Context, id int, offset, limit int) ([]*domain2.User, error) {
	return s.groupRepository.GetUsersByID(ctx, id, offset, limit)
}
func (s *GroupService) GetPostsByID(ctx context.Context, id int, offset, limit int) ([]*domain2.Post, error) {
	return s.groupRepository.GetPostsByID(ctx, id, offset, limit)
}
func (s *GroupService) GetTasksByID(ctx context.Context, id int, offset, limit int) ([]*domain2.Task, error) {
	return s.groupRepository.GetTasksByID(ctx, id, offset, limit)
}
func (s *GroupService) Update(ctx context.Context, id int, group *domain2.Group) (*domain2.Group, error) {
	return s.groupRepository.Update(ctx, id, group)
}
func (s *GroupService) Delete(ctx context.Context, id int) error {
	return s.groupRepository.Delete(ctx, id)
}

func (s *GroupService) StartDiscussion(ctx context.Context, groupID, taskID int) (topics []string, userNames []string, err error) {
	token, ok := ctx.Value("token").(*domain2.Token)
	if !ok {
		return nil, nil, errors.Wrap(domain2.ErrInternal, "token value cannot be converted")
	}

	// 요청한 사용자가 조회해도 되나?
	_, err = s.userService.GetGroupsByID(ctx, token.Subject, 0, 10)
	if err != nil {
		return nil, nil, errors.Wrap(domain2.ErrConstraint, "user cannot access the group")
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

	// bookmark
	gotTask, err := s.taskService.GetByID(ctx, taskID)
	if err != nil {
		return nil, nil, err
	}

	// group bookmark 설정
	_, err = s.groupRepository.Update(ctx, groupID, &domain2.Group{
		Bookmark: gotTask.Range,
	})
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
