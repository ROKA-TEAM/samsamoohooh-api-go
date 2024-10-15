package domain

import (
	"context"
	"time"
)

type UserRoleType string

const (
	UserRoleAdmin UserRoleType = "ADMIN"
	UserRoleGuest UserRoleType = "GUEST"
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
	Create(ctx context.Context, user *User) (*User, error)
	GetByID(ctx context.Context, id int) (*User, error)
	GetGroupsByID(ctx context.Context, id int) ([]*Group, error)
	List(ctx context.Context, limit, offset int) ([]*User, error)
	Update(ctx context.Context, id int, user *User) (*User, error)
	Delete(ctx context.Context, id int) error
}

type UserService interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByID(ctx context.Context, id int) (*User, error)
	GetGroupsByID(ctx context.Context, id int) ([]*Group, error)
	List(ctx context.Context, limit, offset int) ([]*User, error)
	Update(ctx context.Context, id int, user *User) (*User, error)
	Delete(ctx context.Context, id int) error
}
