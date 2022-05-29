package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

const (
	Pass    = "pass"
	Pending = "pending"
	Fail    = "fail"
)

type Flow struct {
	FlowID string `gorm:"primaryKey"`
	//申请人
	ApplicantID string `gorm:"not null"`
	//申请的资源
	ResourceID string `gorm:"not null"`
	//状态
	Statue string `gorm:"not null"`
	//审批人
	ApproverID string
	//申请的App的ID
	AppID string
	//申请的App的名称
	AppName string
	// 配置
	Config datatypes.JSONMap `gorm:"type:json"`
	// 原因
	Reason string
	// 创建时间
	CreatedAt time.Time
	// 更新时间
	UpdatedAt time.Time
}

func (f *Flow) BeforeCreate(_ *gorm.DB) (err error) {
	f.Statue = Pending
	v4, _ := uuid.NewV4()
	f.FlowID = v4.String()
	return
}
