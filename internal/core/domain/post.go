package domain

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title   string
	Content string

	UserID  uint
	GroupID uint
}
