package domain

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	BookTitle   string
	Author      string
	MaxPage     int
	Publisher   string
	Description string
	Bookmark    int

	Users []*User `gorm:"many2many:user_groups;"`
}
