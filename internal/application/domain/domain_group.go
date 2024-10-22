package domain

import (
	"context"
	"time"
)

type Group struct {
	ID          int
	BookTitle   string
	Author      string
	MaxPage     int
	Publisher   string
	Description string
	Bookmark    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time

	// relation
	Users  []*User
	Tasks  []*Task
	Groups []*Group
}

type GroupRepository interface {
	CreateGroup(ctx context.Context, userID int, group *Group) (*Group, error)
	GetGroups(ctx context.Context, offset, limit int) ([]*Group, error)
	GetByGroupID(ctx context.Context, id int) (*Group, error)
	GetUsersByGroupID(ctx context.Context, id int, offset, limit int) ([]*User, error)
	GetPostsByGroupID(ctx context.Context, id int, offset, limit int) ([]*Post, error)
	GetTasksByGroupID(ctx context.Context, id int, offset, limit int) ([]*Task, error)
	UpdateGroup(ctx context.Context, id int, group *Group) (*Group, error)
	DeleteGroup(ctx context.Context, id int) error
	GetUsersLenByGroupID(ctx context.Context, id int) (int, error)
	GetTasksLenByGroupID(ctx context.Context, id int) (int, error)
	AddUser(ctx context.Context, groupID, userID int) error
}

type GroupService interface {
	CreateGroup(ctx context.Context, userID int, group *Group) (*Group, error)
	GetGroups(ctx context.Context, offset, limit int) ([]*Group, error)
	GetByGroupID(ctx context.Context, id int) (*Group, error)
	GetUsersByGroupID(ctx context.Context, id int, offset, limit int) ([]*User, error)
	GetPostsByGroupID(ctx context.Context, id int, offset, limit int) ([]*Post, error)
	GetTasksByGroupID(ctx context.Context, id int, offset, limit int) ([]*Task, error)
	UpdateGroup(ctx context.Context, id int, group *Group) (*Group, error)
	DeleteGroup(ctx context.Context, id int) error
	StartDiscussion(ctx context.Context, groupID, taskID int) (topics []string, userNames []string, err error)

	GroupInviteService
}

type GroupInviteService interface {
	GenerateJoinCode(ctx context.Context, groupID int) (string, error)
	JoinGroupByCode(ctx context.Context, userID int, code string) error
}
