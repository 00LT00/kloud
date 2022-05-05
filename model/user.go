package model

import (
	"gorm.io/gorm"
	"kloud/pkg/util"
	"log"
)

type User struct {
	ID    string `json:"id" gorm:"primaryKey"`
	Name  string `json:"username" gorm:"uniqueIndex;size:40"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Pass  string `json:"password,omitempty"`
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	if u.Pass != "" {
		u.Pass, err = util.PasswordHash(u.Pass)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.Pass, err = util.PasswordHash(u.Pass)
	if err != nil {
		log.Println(err)
		return err
	}
	return
}
