package domain

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	UserID  uint
	TaskID  uint
	Content string
	Feeling string
}
