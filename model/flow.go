package model

import "gorm.io/gorm"

const (
	Pass    = "pass"
	Pending = "pending"
	Fail    = "fail"
)

type Flow struct {
	gorm.Model
	//申请人
	ApplicantID string
	//申请的资源
	ResourceID string
	//状态
	Statue string
	//审批人
	ApproverID string
	//申请的AppID
	AppID string
}

func (f *Flow) BeforeCreate(_ *gorm.DB) (err error) {
	f.Statue = Pending
	return
}
