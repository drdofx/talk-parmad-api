package models

import (
	"time"

	"gorm.io/gorm"
)

type ThreadVote struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ThreadID  uint           `json:"thread_id"`
	UserID    uint           `json:"user_id"`
	Vote      bool           `json:"vote"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
