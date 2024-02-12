package Models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	ID          uint
	Title       string `form:"title" json:"title"  binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
