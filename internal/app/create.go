package app

import (
	"kloud/model"
	"kloud/pkg/DB"
	"log"
)

type Creator interface {
	Create() (model.App, error)
}

type K8sCreator struct {
	*model.Resource
	Config string
	UserID string
}

func (c K8sCreator) Create() (a model.App, err error) {
	log.Println("create k8s app")
	a = model.App{
		ResourceID: c.ResourceID,
		Name:       c.Name,
		Config:     c.Config,
		UserID:     c.UserID,
	}
	db := DB.GetDB()
	err = db.Create(&a).Error
	return
}

type HelmCreator struct {
	*model.Resource
	Config string
	UserID string
}

func (c HelmCreator) Create() (a model.App, err error) {
	log.Println("create helm app")
	a = model.App{
		ResourceID: c.ResourceID,
		Name:       c.Name,
		Config:     c.Config,
		UserID:     c.UserID,
	}
	db := DB.GetDB()
	err = db.Create(&a).Error
	return
}
