package models

import (
	"time"

	"gorm.io/gorm"
)

type Thread struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	Title             string         `json:"title"`
	Text              string         `json:"text" gorm:"type:longtext"`
	ForumID           uint           `json:"forum_id"`
	NumberOfUpvotes   int            `json:"number_of_upvotes"`
	NumberOfDownvotes int            `json:"number_of_downvotes"`
	CreatedBy         uint           `json:"created_by"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}
