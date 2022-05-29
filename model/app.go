package model

import (
	"gorm.io/datatypes"
	"time"
)

type App struct {
	AppID      string            `gorm:"primaryKey"`
	UserID     string            `gorm:"not null"`
	ResourceID string            `gorm:"not null"`
	Name       string            `gorm:"not null"`
	Config     datatypes.JSONMap `gorm:"type:json"`
	// 创建时间
	CreatedAt time.Time
	// 更新时间
	UpdatedAt time.Time
}
