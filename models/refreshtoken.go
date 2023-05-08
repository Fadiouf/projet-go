package models

import "time"

// RefreshToken model
type RefreshToken struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Token     string    `json:"token" gorm:"not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
}
