package user

import (
	"kloud/model"
	"kloud/pkg/DB"
)

func isExist(id string) bool {
	db := DB.GetDB()
	u := new(model.User)
	u.ID = id
	var cnt int64
	db.Model(&model.User{}).Where(u).Count(&cnt)
	return cnt > 0
}
