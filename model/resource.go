package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Resource struct {
	//自动生成
	ResourceID string `gorm:"primaryKey"`
	Name       string
	Folder     string
}

func (r *Resource) BeforeCreate(_ *gorm.DB) (err error) {
	v4, _ := uuid.NewV4()
	r.ResourceID = v4.String()
	return
}
