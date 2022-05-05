package casbin

import (
	"errors"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"kloud/pkg/DB"
	"kloud/pkg/conf"
	"log"
)

const (
	Super    = "root"
	Admin    = "admin"
	Approve  = "distribute"
	Approver = "distributor"
	Importer = "importer"
	Import   = "import"
	Any      = "*"
)

var e *enforcer

func init() {
	a, err := gormadapter.NewAdapterByDB(DB.GetDB())
	if err != nil {
		panic(err)
	}
	syncedEnforcer, err := casbin.NewSyncedEnforcer(conf.GetConf().Casbin.Model(), a)
	if err != nil {
		panic(err)
	}
	e = (*enforcer)(syncedEnforcer)
	if len(e.GetPolicy()) < 3 {
		ok, err := e.AddPolicies([][]string{
			{Approver, Any, Approve},
			{Importer, Any, Import},
			{Super, Any, Any},
		})
		if !ok || err != nil {
			if err != nil {
				log.Println(err.Error())
			}
			panic(errors.New("add default policy error"))
		}
		ok, err = e.AddRolesForUser(Admin, []string{Importer, Approver})
		if !ok || err != nil {
			if err != nil {
				log.Println(err.Error())
			}
			panic(errors.New("add default policy error"))
		}
	}
}

func GetEnforcer() *enforcer {
	return e
}
