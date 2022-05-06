package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

const (
	K8s  = "k8s"
	Helm = "helm"
)

type Resource struct {
	//自动生成
	ResourceID string `json:"resource_id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Folder     string `json:"folder"`
	Type       string `json:"type"`
}

func (r *Resource) BeforeCreate(_ *gorm.DB) (err error) {
	v4, _ := uuid.NewV4()
	r.ResourceID = v4.String()
	return
}
