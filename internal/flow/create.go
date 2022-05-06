package flow

import (
	"github.com/gin-gonic/gin"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/util"
	"log"
	"net/http"
)

func RestCreate(c *gin.Context) {
	id := c.PostForm("resource_id")
	config := c.PostForm("config")
	if id == "" {
		c.JSON(util.MakeResp(http.StatusBadRequest, 0, "resource_id null"))
		return
	}
	v, _ := c.Get("user")
	u := v.(model.User)
	err := createFlow(u.ID, id, config)
	if err != nil {
		c.JSON(util.MakeResp(http.StatusBadRequest, 1, err.Error()))
		return
	}
	c.JSON(util.MakeOkResp("create flow success"))
}

func createFlow(applicantID, resourceID, config string) error {
	f := new(model.Flow)
	f.ResourceID = resourceID
	f.ApplicantID = applicantID
	f.Config = config
	db := DB.GetDB()
	err := db.Create(f).Error
	if err != nil {
		log.Println(err.Error())
	}
	return err
}
