package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type App struct {
	gorm.Model
	AppID      string `gorm:"uniqueIndex;size:40"`
	ResourceID string
	Name       string
	Remarks    string
}

func (a *App) BeforeCreate(_ *gorm.DB) (err error) {
	v4, _ := uuid.NewV4()
	a.AppID = v4.String()
	return
}
