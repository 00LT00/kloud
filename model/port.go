package model

type PortMapping struct {
	Port       int    `gorm:"primaryKey"`
	TargetPort int    `gorm:"not null"`
	AppID      string `gorm:"not null"`
}
