package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/casbin"
	"kloud/pkg/util"
	"net/http"
)

func RestRegister(c *gin.Context) {
	u := new(model.User)
	if err := c.ShouldBindJSON(u); err != nil {
		c.JSON(util.MakeResp(http.StatusBadRequest, 1, err.Error()))
		return
	}
	if u.Name == "" || u.Pass == "" {
		c.JSON(util.MakeResp(http.StatusBadRequest, 0, "username or password null"))
		return
	}

	//生成id
	v4, _ := uuid.NewV4()
	u.ID = v4.String()
	db := DB.GetDB().Begin()
	if err := db.Create(u).Error; err != nil {
		db.Rollback()
		c.JSON(util.MakeResp(http.StatusBadRequest, 1, err.Error()))
		return
	}

	//判断是否要初始化
	e := casbin.GetEnforcer()
	if len(e.GetSuperUsers()) == 0 {
		ok := e.SetSupper(u.ID)
		if !ok {
			db.Rollback()
			c.JSON(util.MakeResp(http.StatusBadRequest, 0, "init super admin error"))
			return
		}
	}

	db.Commit()
	c.JSON(util.MakeOkResp("register success"))
}
