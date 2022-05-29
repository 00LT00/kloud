package model

import (
	"errors"
	"gorm.io/gorm"
)

type PortMapping struct {
	Port       int    `gorm:"primaryKey"`
	TargetPort int    `json:"targetPort" gorm:"not null"`
	AppID      string `gorm:"not null"`
}

func (p PortMapping) BeforeCreate(_ *gorm.DB) (err error) {
	//校验端口范围合法
	if p.Port < 1 || p.Port > 65535 {
		return errors.New("port out of range")
	}
	if p.TargetPort < 1 || p.TargetPort > 65535 {
		return errors.New("targetPort out of range")
	}
	return
}

func (p PortMapping) BeforeUpdate(_ *gorm.DB) (err error) {
	//校验端口范围合法
	if p.Port < 1 || p.Port > 65535 {
		return errors.New("port out of range")
	}
	if p.TargetPort < 1 || p.TargetPort > 65535 {
		return errors.New("targetPort out of range")
	}
	return
}
