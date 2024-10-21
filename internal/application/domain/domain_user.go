package domain

import (
	"context"
	"time"
)

type UserRoleType string

const (
	UserRoleAdmin   UserRoleType = "ADMIN"
	UserRoleMANAGER UserRoleType = "MANAGER"

	UserRoleUser UserRoleType = "USER"
)

type UserSocialType string

const (
	UserSocialGoogle UserSocialType = "GOOGLE"
	UserSocialApple  UserSocialType = "APPLE"
	UserSocialKaKao  UserSocialType = "KAKAO"
)

type User struct {
	ID         int
	Name       string
	Resolution string
	Role       UserRoleType
	Social     UserSocialType
	SocialSub  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time

	// relation
	Groups   []*Group
	Topics   []*Topic
	Posts    []*Post
	Comments []*Comment
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetByUserID(ctx context.Context, id int) (*User, error)
	GetByUserSub(ctx context.Context, sub string) (*User, error)
	GetGroupsByUserID(ctx context.Context, id int, limit, offset int) ([]*Group, error)
	GetUsers(ctx context.Context, limit, offset int) ([]*User, error)
	UpdateUser(ctx context.Context, id int, user *User) (*User, error)
	DeleteUser(ctx context.Context, id int) error

	IsUserInGroup(ctx context.Context, userID, groupID int) (bool, error)
}

type UserService interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetByUserID(ctx context.Context, id int) (*User, error)
	GetByUserSub(ctx context.Context, sub string) (*User, error)
	GetGroupsByUserID(ctx context.Context, id int, limit, offset int) ([]*Group, error)
	GetUsers(ctx context.Context, limit, offset int) ([]*User, error)
	UpdateUser(ctx context.Context, id int, user *User) (*User, error)
	DeleteUser(ctx context.Context, id int) error

	IsUserInGroup(ctx context.Context, userID, groupID int) (bool, error)
}
