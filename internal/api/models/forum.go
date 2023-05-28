package models

import (
	"time"

	"gorm.io/gorm"
)

type Forum struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	ForumName        string         `json:"forum_name" gorm:"unique;type:varchar(255)"`
	IntroductionText string         `json:"introduction_text" gorm:"type:text"`
	ForumImage       *string        `json:"forum_image"`
	Category         *string        `json:"category"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
}
