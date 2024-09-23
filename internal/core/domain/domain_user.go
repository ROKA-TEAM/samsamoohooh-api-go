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
	Name       string     `gorm:"not null;unique"`
	Resolution string     `gorm:"not null"`
	Role       RoleType   `gorm:"not null"`
	Sub        string     `gorm:"not null"`
	Social     SocialType `gorm:"not null"`

	Groups []*Group `gorm:"many2many:user_groups;"`

	Posts    []Post    `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:UserID"`
	Subjects []Subject `gorm:"foreignKey:UserID"`
}
