package model

type PortMapping struct {
	Port       int    `gorm:"primaryKey"`
	TargetPort int    `json:"targetPort" gorm:"not null"`
	AppID      string `gorm:"not null"`
}
