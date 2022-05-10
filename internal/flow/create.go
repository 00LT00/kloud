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
	type req struct {
		ResourceID string `json:"resource_id"`
		Config     string `json:"config"`
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
	v, _ := c.Get("user")
	u := v.(model.User)
	err = createFlow(u.ID, r.ResourceID, r.Config)
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
