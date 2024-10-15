package domain

import "time"

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
	SocialType UserSocialType
	SocialSub  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
