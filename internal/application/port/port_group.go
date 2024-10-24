package port

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
)

type GroupRepository interface {
	CreateGroup(ctx context.Context, userID int, group *domain.Group) (*domain.Group, error)
	GetGroups(ctx context.Context, offset, limit int) ([]*domain.Group, error)
	GetByGroupID(ctx context.Context, id int) (*domain.Group, error)
	GetUsersByGroupID(ctx context.Context, id int, offset, limit int) ([]*domain.User, error)
	GetPostsByGroupID(ctx context.Context, id int, offset, limit int) ([]*domain.Post, error)
	GetTasksByGroupID(ctx context.Context, id int, offset, limit int) ([]*domain.Task, error)
	UpdateGroup(ctx context.Context, id int, group *domain.Group) (*domain.Group, error)
	DeleteGroup(ctx context.Context, id int) error
	GetUsersLenByGroupID(ctx context.Context, id int) (int, error)
	GetTasksLenByGroupID(ctx context.Context, id int) (int, error)
	AddUser(ctx context.Context, groupID, userID int) error
	RemoveUser(ctx context.Context, groupID, userID int) error
}

type GroupService interface {
	CreateGroup(ctx context.Context, userID int, group *domain.Group) (*domain.Group, error)
	GetGroups(ctx context.Context, offset, limit int) ([]*domain.Group, error)
	GetByGroupID(ctx context.Context, id int) (*domain.Group, error)
	GetUsersByGroupID(ctx context.Context, id int, offset, limit int) ([]*domain.User, error)
	GetPostsByGroupID(ctx context.Context, id int, offset, limit int) ([]*domain.Post, error)
	GetTasksByGroupID(ctx context.Context, id int, offset, limit int) ([]*domain.Task, error)
	UpdateGroup(ctx context.Context, id int, group *domain.Group) (*domain.Group, error)
	DeleteGroup(ctx context.Context, id int) error
	StartDiscussion(ctx context.Context, groupID, taskID int) (topics []string, userNames []string, err error)
	LeaveGroup(ctx context.Context, groupID, userID int) error

	GenerateJoinCode(ctx context.Context, groupID int) (string, error)
	JoinGroupByCode(ctx context.Context, userID int, code string) error
}
