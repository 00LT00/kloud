package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/util"
	"log"
	"net/http"
)

func RestGetByUser(c *gin.Context) {
	v, _ := c.Get("user")
	u := v.(model.User)
	db := DB.GetDB()
	var apps []*model.App
	err := db.Where(&model.App{UserID: u.ID}).Find(&apps).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)
		c.JSON(util.MakeResp(http.StatusInternalServerError, 0, "unknown error"))
		return
	}
	c.JSON(util.MakeOkResp(apps))
}

func RestGetAll(c *gin.Context) {
	db := DB.GetDB()
	var apps []*model.App
	err := db.Find(&apps).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)
		c.JSON(util.MakeResp(http.StatusInternalServerError, 0, "unknown error"))
		return
	}
	c.JSON(util.MakeOkResp(apps))
}
