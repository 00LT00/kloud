package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

const (
	K8s  = "k8s"
	Helm = "helm"
)

type Resource struct {
	//自动生成
	ResourceID string `json:"resource_id" gorm:"primaryKey"`
	Name       string `json:"name" gorm:"not null"`
	Folder     string `json:"folder" gorm:"not null"`
	Type       string `json:"type" gorm:"not null"`
	MaxNum     int    `json:"max_num" gorm:"not null"`
	// 创建时间
	CreatedAt time.Time
	// 更新时间
	UpdatedAt time.Time
}

func (r *Resource) BeforeCreate(_ *gorm.DB) (err error) {
	v4, _ := uuid.NewV4()
	r.ResourceID = v4.String()
	return
}

func (r *Resource) GetTemplateFilename() string {
	return r.Folder + "template.yaml"
}

func (r *Resource) GetConfigFilename() string {
	return r.Folder + "config.json"
}
