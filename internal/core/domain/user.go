package domain

import "gorm.io/gorm"

type RoleType string

const (
	Admin  RoleType = "ADMIN"
	Member RoleType = "MEMBER"
	Guest  RoleType = "GUEST"
)

type SocialType string

const (
	Google SocialType = "GOOGLE"
	Kakao  SocialType = "KAKAO"
	Apple  SocialType = "APPLE"
)

type User struct {
	gorm.Model
	Name       string `gorm:"unique"`
	Resolution string
	Role       RoleType
	Sub        string
	Social     SocialType

	Groups []*Group `gorm:"many2many:user_groups;"`

	Posts []Post `gorm:"foreignKey:UserID"`
}
