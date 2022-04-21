package user

import (
	"github.com/gin-gonic/gin"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/casbin"
	"kloud/pkg/util"
	"net/http"
)

type restUsers struct {
	Users []string `json:"users"`
}

func RestAddAdmin(c *gin.Context) {
	req := new(restUsers)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(util.MakeResp(http.StatusBadRequest, 1, err.Error()))
		return
	}
	db := DB.GetDB()
	e := casbin.GetEnforcer()
	successes := make([]string, 0)
	for _, id := range req.Users {
		u := new(model.User)
		u.ID = id
		var cnt int64
		db.Model(&model.User{}).Where(u).Count(&cnt)
		if cnt > 0 {
			if e.AddAdmin(id) {
				successes = append(successes, id)
			}
		}
	}
	c.JSON(util.MakeOkResp(successes))
}
