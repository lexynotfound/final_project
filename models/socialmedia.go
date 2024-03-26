package models

import "time"

type SocialMedia struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	URL       string    `gorm:"type:text;not null"`
	UserID    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"default:null"`
}
