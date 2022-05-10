package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type App struct {
	AppID      string `gorm:"primaryKey"`
	UserID     string `gorm:"not null"`
	ResourceID string `gorm:"not null"`
	Name       string `gorm:"not null"`
	Config     string
	// 创建时间
	CreatedAt time.Time
	// 更新时间
	UpdatedAt time.Time
}

func (a *App) BeforeCreate(_ *gorm.DB) (err error) {
	v4, _ := uuid.NewV4()
	a.AppID = v4.String()
	return
}
