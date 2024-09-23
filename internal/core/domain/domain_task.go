package domain

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	GroupID  uint
	Deadline time.Time
	Range    int

	Subjects []Subject `gorm:"foreignKey:TaskID"`
}
