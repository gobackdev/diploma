package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"index;not null"` // внешний ключ на пользователя
	OrderNumber string `gorm:"index;not null;unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
