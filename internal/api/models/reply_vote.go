package models

import (
	"time"

	"gorm.io/gorm"
)

type ReplyVote struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ReplyID   uint           `json:"reply_id"`
	UserID    uint           `json:"user_id"`
	Vote      bool           `json:"vote"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
