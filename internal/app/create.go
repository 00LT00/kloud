package app

import (
	"kloud/model"
	"kloud/pkg/DB"
	"log"
)

type Creator interface {
	Create(app *model.App) error
}

type HelmCreator struct {
	*model.Resource
}

func (c HelmCreator) Create(a *model.App) (err error) {
	log.Println("create helm app")
	a = &model.App{
		ResourceID: c.ResourceID,
		Name:       c.Name,
	}
	db := DB.GetDB()
	err = db.Create(&a).Error
	return
}
