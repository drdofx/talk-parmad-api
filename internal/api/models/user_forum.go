package models

import (
	"time"

	"gorm.io/gorm"
)

type UserForum struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id"`
	ForumID   uint           `json:"forum_id"`
	IsRemoved bool           `json:"is_removed" default:"false"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
