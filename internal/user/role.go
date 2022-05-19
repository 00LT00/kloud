package user

import (
	"github.com/gin-gonic/gin"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/casbin"
	"kloud/pkg/util"
	"net/http"
)

func RestAddAdmin(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(util.MakeResp(http.StatusBadRequest, 0, "id none"))
		return
	}
	e := casbin.GetEnforcer()
	if !isExist(id) {
		c.JSON(util.MakeResp(http.StatusNotFound, 0, "user none"))
	}
	if !e.AddAdmin(id) {
		c.JSON(util.MakeResp(http.StatusInternalServerError, 0, "add admin role error"))
	}
	c.JSON(util.MakeOkResp("success"))
}

func RestDeleteAdmin(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(util.MakeResp(http.StatusBadRequest, 0, "id none"))
		return
	}
	e := casbin.GetEnforcer()
	if !isExist(id) {
		c.JSON(util.MakeResp(http.StatusNotFound, 0, "user none"))
	}
	if !e.DeleteAdmin(id) {
		c.JSON(util.MakeResp(http.StatusInternalServerError, 0, "delete admin role error"))
	}
	c.JSON(util.MakeOkResp("success"))
}

func RestGetAdmin(c *gin.Context) {
	e := casbin.GetEnforcer()
	ids := e.GetAdminUsers()
	db := DB.GetDB()
	users := make([]*model.User, 0, len(ids))
	for _, id := range ids {
		user := new(model.User)
		db.Where(&model.User{ID: id}).First(user)
		if user != nil {
			user.Pass = ""
		}
		users = append(users, user)
	}
	c.JSON(util.MakeOkResp(users))
}

func getRole(id string) string {
	e := casbin.GetEnforcer()
	if ok, _ := e.HasRoleForUser(id, casbin.Super); ok {
		return casbin.Super
	}
	if ok, _ := e.HasRoleForUser(id, casbin.Admin); ok {
		return casbin.Admin
	}
	return casbin.User
}
