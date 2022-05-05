package casbin

import (
	"github.com/casbin/casbin/v2"
	"log"
)

type enforcer casbin.SyncedEnforcer

func (e *enforcer) AddAdmin(id string) bool {
	ok, err := e.HasRoleForUser(id, Admin)
	if err != nil {
		log.Println(err.Error())
	}
	if ok {
		return ok
	}
	ok, err = e.AddRoleForUser(id, Admin)
	if err != nil {
		log.Println(err.Error())
	}
	return ok
}

func (e *enforcer) DeleteAdmin(id string) bool {
	ok, err := e.DeleteRoleForUser(id, Admin)
	if err != nil {
		log.Println(err.Error())
	}
	return ok
}

func (e *enforcer) GetAdminUsers() []string {
	users, err := e.GetUsersForRole(Admin)
	if err != nil {
		log.Println(err.Error())
	}
	return users
}

func (e *enforcer) GetSuperUsers() []string {
	users, err := e.GetUsersForRole(Super)
	if err != nil {
		log.Println(err.Error())
	}
	return users
}

func (e *enforcer) SetSupper(id string) bool {
	ok, err := e.AddRoleForUser(id, Super)
	if err != nil {
		log.Println(err.Error())
	}
	return ok
}
