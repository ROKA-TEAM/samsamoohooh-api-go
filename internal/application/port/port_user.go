package port

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetByUserID(ctx context.Context, id int) (*domain.User, error)
	GetByUserSub(ctx context.Context, sub string) (*domain.User, error)
	GetGroupsByUserID(ctx context.Context, id int, limit, offset int) ([]*domain.Group, error)
	GetUsers(ctx context.Context, limit, offset int) ([]*domain.User, error)
	UpdateUser(ctx context.Context, id int, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	IsUserInGroup(ctx context.Context, userID, groupID int) (bool, error)
}

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetByUserID(ctx context.Context, id int) (*domain.User, error)
	GetByUserSub(ctx context.Context, sub string) (*domain.User, error)
	GetGroupsByUserID(ctx context.Context, id int, limit, offset int) ([]*domain.Group, error)
	GetUsers(ctx context.Context, limit, offset int) ([]*domain.User, error)
	UpdateUser(ctx context.Context, id int, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	IsUserInGroup(ctx context.Context, userID, groupID int) (bool, error)
}
