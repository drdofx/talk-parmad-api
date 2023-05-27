package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"type:varchar(255)"`
	Email        string         `json:"email" gorm:"unique;type:varchar(255)"`
	Password     string         `json:"-" gorm:"type:varchar(255)"`
	Role         string         `json:"role" gorm:"type:ENUM('Admin', 'User');default:'User'"`
	ProfileImage *string        `json:"profile_image" gorm:"type:varchar(255)"`
	NIM          *string        `json:"nim,omitempty" gorm:"unique;type:varchar(255)"`
	Status       *string        `json:"status" gorm:"type:ENUM('Active', 'Inactive');default:'Active'"`
	Prodi        string         `json:"prodi" gorm:"type:varchar(255)"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
