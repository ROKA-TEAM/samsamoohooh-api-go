package service

import (
	"context"

	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/infra/exception"
	"time"

	"github.com/google/uuid"
)

var _ port.GroupService = (*GroupService)(nil)

const (
	JoinCodeExpireTime = time.Second * 60 * 24 * 2 // 2day
)

type GroupService struct {
	groupRepository port.GroupRepository
	userService     port.UserService
	taskService     port.TaskService
	redisRepository port.RedisRepository
}

func NewGroupService(
	groupRepository port.GroupRepository,
	userService port.UserService,
	taskService port.TaskService,
	keyValueRepository port.RedisRepository,
) *GroupService {
	return &GroupService{
		groupRepository: groupRepository,
		userService:     userService,
		taskService:     taskService,
		redisRepository: keyValueRepository,
	}
}

func (s *GroupService) CreateGroup(ctx context.Context, userID int, group *domain.Group) (*domain.Group, error) {
	createdUser, err := s.groupRepository.CreateGroup(ctx, userID, group)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *GroupService) GetGroups(ctx context.Context, offset, limit int) ([]*domain.Group, error) {
	listGroup, err := s.groupRepository.GetGroups(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	return listGroup, nil
}

func (s *GroupService) GetByGroupID(ctx context.Context, id int) (*domain.Group, error) {
	gotGroup, err := s.groupRepository.GetByGroupID(ctx, id)
	if err != nil {
		return nil, err
	}

	return gotGroup, nil
}

func (s *GroupService) GetUsersByGroupID(ctx context.Context, id int, offset, limit int) ([]*domain.User, error) {
	listUser, err := s.groupRepository.GetUsersByGroupID(ctx, id, offset, limit)
	if err != nil {
		return nil, err
	}

	return listUser, nil
}

func (s *GroupService) GetPostsByGroupID(ctx context.Context, id int, offset, limit int) ([]*domain.Post, error) {
	listGroup, err := s.groupRepository.GetPostsByGroupID(ctx, id, offset, limit)
	if err != nil {
		return nil, err
	}

	return listGroup, nil
}

func (s *GroupService) GetTasksByGroupID(ctx context.Context, id int, offset, limit int) ([]*domain.Task, error) {
	listTask, err := s.groupRepository.GetTasksByGroupID(ctx, id, offset, limit)
	if err != nil {
		return nil, err
	}

	return listTask, nil
}

func (s *GroupService) UpdateGroup(ctx context.Context, id int, group *domain.Group) (*domain.Group, error) {
	updatedGroup, err := s.groupRepository.UpdateGroup(ctx, id, group)
	if err != nil {
		return nil, err
	}

	return updatedGroup, nil
}

func (s *GroupService) DeleteGroup(ctx context.Context, id int) error {
	err := s.groupRepository.DeleteGroup(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *GroupService) StartDiscussion(ctx context.Context, groupID, taskID int) (topics []string, userNames []string, err error) {

	topicsLen, err := s.taskService.GetTopicsLenByTaskID(ctx, taskID)
	if err != nil {
		return nil, nil, err
	}

	listTopic, err := s.taskService.GetTopicsByTaskID(ctx, taskID, 0, topicsLen)
	if err != nil {
		return nil, nil, err
	}

	usersLen, err := s.groupRepository.GetUsersLenByGroupID(ctx, groupID)
	if err != nil {
		return nil, nil, err
	}
	listUser, err := s.groupRepository.GetUsersByGroupID(ctx, groupID, 0, usersLen)
	if err != nil {
		return nil, nil, err
	}

	// bookmark
	gotTask, err := s.taskService.GetByTaskID(ctx, taskID)
	if err != nil {
		return nil, nil, err
	}

	// group bookmark 설정
	_, err = s.groupRepository.UpdateGroup(ctx, groupID, &domain.Group{
		Bookmark: gotTask.Range,
	})
	if err != nil {
		return nil, nil, err
	}

	for _, topic := range listTopic {
		topics = append(topics, topic.Topic)
	}

	for _, user := range listUser {
		userNames = append(userNames, user.Name)
	}

	return topics, userNames, nil
}

func (s *GroupService) GenerateJoinCode(ctx context.Context, groupID int) (string, error) {
	joinCode := uuid.New().String()

	err := s.redisRepository.Set(ctx, joinCode, groupID, JoinCodeExpireTime)
	if err != nil {
		return "", err
	}

	return joinCode, nil
}

func (s *GroupService) JoinGroupByCode(ctx context.Context, userID int, code string) error {
	groupID, err := s.redisRepository.GetInt(ctx, code)
	if err != nil {
		return err
	}

	// 이미 참가한 사용자인지 확인
	isIn, err := s.userService.IsUserInGroup(ctx, userID, groupID)
	if err != nil {
		return err
	}

	if isIn {
		return exception.NewWithoutErr(
			exception.ErrBizGroupAlreadyJoined,
			exception.StatusForbidden,
			"already joined",
		)
	}

	err = s.groupRepository.AddUser(ctx, groupID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *GroupService) LeaveGroup(ctx context.Context, userID, groupID int) error {
	err := s.groupRepository.RemoveUser(ctx, groupID, userID)
	if err != nil {
		return err
	}

	usersLen, err := s.groupRepository.GetUsersLenByGroupID(ctx, groupID)
	if err != nil {
		return err
	}

	// 모임을 삭제하는 경우
	if usersLen == 0 {
		err := s.groupRepository.DeleteGroup(ctx, groupID)
		if err != nil {
			return err
		}
	}

	return nil
}
