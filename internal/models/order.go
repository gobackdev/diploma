package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"index;not null"`
	OrderNumber string `gorm:"size:64;not null;unique"`
	Status      string `gorm:"size:16;not null;default:NEW"`
	Accrual     *int
	UploadedAt  time.Time `gorm:"autoCreateTime"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
