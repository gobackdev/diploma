package models

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Login        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	CreatedAt    time.Time
}
