package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type App struct {
	AppID      string `gorm:"primaryKey"`
	UserID     string
	ResourceID string
	Name       string
	Config     string
}

func (a *App) BeforeCreate(_ *gorm.DB) (err error) {
	v4, _ := uuid.NewV4()
	a.AppID = v4.String()
	return
}
