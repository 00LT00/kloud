package flow

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kloud/internal/app"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/util"
	"log"
	"net/http"
)

func RestApprove(c *gin.Context) {
	id := c.Param("id")
	type reqApprove struct {
		Reason string
		Status string
	}
	a := new(reqApprove)
	_ = c.ShouldBindJSON(a)
	switch a.Status {
	case model.Pass, model.Fail, "":
	default:
		c.JSON(util.MakeResp(http.StatusInternalServerError, 0, "status error"))
		return
	}
	db := DB.GetDB()
	f := new(model.Flow)
	f.FlowID = id
	err := db.Where(f).First(f).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(util.MakeResp(http.StatusNotFound, 0, "flow none"))
			return
		}
		log.Println(err)
		c.JSON(util.MakeResp(http.StatusInternalServerError, 0, "unknown error"))
		return
	}
	f.Reason = a.Reason
	f.Statue = a.Status
	v, _ := c.Get("user")
	u := v.(model.User)
	f.ApproverID = u.ID
	err = approve(f)
	if err != nil {
		log.Println(err)
		c.JSON(util.MakeResp(http.StatusInternalServerError, 1, err.Error()))
		return
	}
	c.JSON(util.MakeOkResp("approve success"))
}

func approve(f *model.Flow) error {
	db := DB.GetDB()
	if f.Statue == model.Fail {
		err := db.Save(f).Error
		return err
	}
	r := new(model.Resource)
	r.ResourceID = f.ResourceID
	db.First(r)
	var c app.Creator
	switch r.Type {
	case model.K8s:
		c = &app.K8sCreator{Resource: r, Config: f.Config, UserID: f.ApplicantID}
	case model.Helm:
		c = &app.HelmCreator{Resource: r, Config: f.Config, UserID: f.ApplicantID}
	}
	a, err := c.Create()
	if err != nil {
		f.Statue = model.Fail
		f.Reason = fmt.Sprintf("create error, reason:%s", err.Error())
	} else {
		f.AppID = a.AppID
	}
	err = db.Save(f).Error
	return err
}
