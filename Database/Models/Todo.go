package Models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	ID          uint
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
