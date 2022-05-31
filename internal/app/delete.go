package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/k8s"
	"kloud/pkg/util"
	"net/http"
)

func RestDelete(c *gin.Context) {
	id := c.Param("id")
	v, _ := c.Get("user")
	u := v.(model.User)
	db := DB.GetDB()
	var cnt int64
	db.Model(&model.App{}).Where(&model.App{AppID: id, UserID: u.ID}).Count(&cnt)
	if cnt == 0 {
		c.JSON(util.MakeResp(http.StatusNotFound, 0, "app none"))
		return
	}
	err := deleteApp(id)
	if err != nil {
		c.JSON(util.MakeResp(http.StatusInternalServerError, 1, err.Error()))
		return
	}
	c.JSON(util.MakeOkResp("delete app success"))
}

func deleteApp(id string) (err error) {
	db := DB.GetDB()
	var cnt int64
	db.Model(&model.App{}).Where(&model.App{AppID: id}).Count(&cnt)
	if cnt == 0 {
		return errors.New("app none")
	}
	var ports []int
	if err = db.Model(&model.PortMapping{}).Where(&model.PortMapping{AppID: id}).Pluck("port", &ports).Error; err != nil {
		return
	}
	tx := db.Begin()
	err = tx.Where(&model.PortMapping{AppID: id}).Delete(&model.PortMapping{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Where(&model.App{AppID: id}).Delete(&model.App{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, port := range ports {
		k8s.StopPort(port)
	}
	err = k8s.DeleteApp(id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return
}
