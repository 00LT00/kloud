package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

const (
	Pass    = "pass"
	Pending = "pending"
	Fail    = "fail"
)

type Flow struct {
	FlowID string `gorm:"primaryKey"`
	//申请人
	ApplicantID string
	//申请的资源
	ResourceID string
	//状态
	Statue string
	//审批人
	ApproverID string
	//申请的App的ID
	AppID string
	// 配置
	Config string
	// 原因
	Reason string
}

func (f *Flow) BeforeCreate(_ *gorm.DB) (err error) {
	f.Statue = Pending
	v4, _ := uuid.NewV4()
	f.FlowID = v4.String()
	return
}
