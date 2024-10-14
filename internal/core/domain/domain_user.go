package domain

import "time"

type RoleType string

const (
	RoleAdmin  RoleType = "ADMIN"
	RoleMember RoleType = "MEMBER"
	RoleGuest  RoleType = "GUEST"
)

type SocialType string

const (
	SocialGoogle SocialType = "GOOGLE"
	SocialKakao  SocialType = "KAKAO"
	SocialApple  SocialType = "APPLE"
)

type User struct {
	ID         int
	Name       string
	Resolution string
	Role       RoleType
	Social     SocialType

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Groups   []Group
	Posts    []Post
	Comments []Comment
	Subjects []Subject
}

// domain 정의 -> router, dto 정의 -> Repository 정의 -> service 정의 -> handler 구현
