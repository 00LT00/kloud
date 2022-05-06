package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/util"
	"net/http"
)

func RestGetAllUser(c *gin.Context) {
	db := DB.GetDB()
	var users []*model.User
	err := db.Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(util.MakeResp(http.StatusInternalServerError, 1, err.Error()))
		return
	}
	for i := range users {
		users[i].Pass = ""
	}
	c.JSON(util.MakeOkResp(users))
}
