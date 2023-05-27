package models

import (
	"time"

	"gorm.io/gorm"
)

type Moderator struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Nickname  *string        `json:"nickname"`
	Rank      string         `json:"rank" gorm:"type:ENUM('Head', 'Member');default:'Member'"`
	UserID    uint           `json:"user_id"`
	ForumID   uint           `json:"forum_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
