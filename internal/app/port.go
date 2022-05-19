package app

import (
	"kloud/model"
	"kloud/pkg/DB"
)

func getPortMapping(id string) []model.PortMapping {
	db := DB.GetDB()
	var portMappings []model.PortMapping
	db.Where(&model.PortMapping{AppID: id}).Find(&portMappings)
	return portMappings
}
