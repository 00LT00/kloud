package user

import (
	"github.com/gin-gonic/gin"
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
	c.JSON(util.MakeOkResp(ids))
}
