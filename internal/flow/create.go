package flow

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/util"
	"log"
	"net/http"
)

func RestCreate(c *gin.Context) {
	type req struct {
		ResourceID string            `json:"resource_id"`
		Name       string            `json:"name"`
		Config     datatypes.JSONMap `json:"config"`
	}
	r := new(req)
	err := c.ShouldBindJSON(r)
	if err != nil {
		log.Println(err.Error())
	}
	if r.ResourceID == "" {
		c.JSON(util.MakeResp(http.StatusBadRequest, 0, "resource_id null"))
		return
	}
	if r.Name == "" {
		c.JSON(util.MakeResp(http.StatusBadRequest, 0, "name null"))
		return
	}
	v, _ := c.Get("user")
	u := v.(model.User)
	err = createFlow(u.ID, r.Name, r.ResourceID, r.Config)
	if err != nil {
		c.JSON(util.MakeResp(http.StatusBadRequest, 1, err.Error()))
		return
	}
	c.JSON(util.MakeOkResp("create flow success"))
}

func createFlow(applicantID, appName, resourceID string, config datatypes.JSONMap) error {
	db := DB.GetDB()
	r := new(model.Resource)
	r.ResourceID = resourceID
	err := db.Where(r).First(r).Error
	if err != nil {
		return err
	}
	var cnt int64
	db.Model(&model.App{}).Where(&model.App{ResourceID: resourceID, UserID: applicantID}).Count(&cnt)
	sum := cnt
	db.Model(&model.Flow{}).Where(&model.Flow{ResourceID: resourceID, ApplicantID: applicantID, Statue: model.Pending}).Count(&cnt)
	sum += cnt
	if sum >= int64(r.MaxNum) {
		return errors.New("out of max num")
	}
	f := &model.Flow{
		ResourceID:  resourceID,
		AppName:     appName,
		ApplicantID: applicantID,
		Config:      config,
	}
	err = db.Create(f).Error
	if err != nil {
		log.Println(err.Error())
	}
	return err
}
