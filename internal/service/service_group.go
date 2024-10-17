package service

import (
	"context"
	"github.com/pkg/errors"
	"samsamoohooh-go-api/internal/domain"
)

var _ domain.GroupService = &GroupService{}

type GroupService struct {
	groupRepository domain.GroupRepository
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
