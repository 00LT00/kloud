package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"kloud/pkg/conf"
	"kloud/pkg/database"
)

var e *casbin.SyncedEnforcer

func init() {
	a, err := gormadapter.NewAdapterByDB(database.GetDB())
	if err != nil {
		panic(err)
	}
	e, err = casbin.NewSyncedEnforcer(conf.GetConf().Casbin.Model(), a)
	if err != nil {
		panic(err)
	}
}

func GetEnforcer() *casbin.SyncedEnforcer {
	return e
}
