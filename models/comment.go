package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	Message   string    `gorm:"type:text;not null"`
	UserID    uint      `gorm:"not null"`
	PhotoID   uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"default:null"`
}
