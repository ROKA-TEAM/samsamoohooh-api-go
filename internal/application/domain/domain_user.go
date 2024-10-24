package domain

import (
	"time"
)

type UserRoleType string

const (
	UserRoleAdmin   UserRoleType = "ADMIN"
	UserRoleManager UserRoleType = "MANAGER"
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
