package models

import "time"

type Photo struct {
    ID        uint      `gorm:"primaryKey"`
    Title     string    `gorm:"not null"`
    URL       string    `gorm:"type:text;not null"`
    Caption   string    `gorm:"type:text"`
    UserID    uint      `gorm:"not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
    DeletedAt time.Time `gorm:"default:null"`
}