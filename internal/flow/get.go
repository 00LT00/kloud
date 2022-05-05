package flow

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/util"
	"log"
	"net/http"
	"strconv"
)

func RestGetPending(c *gin.Context) {
	db := DB.GetDB()
	var flows []*model.Flow
	err := db.Where(&model.Flow{Statue: model.Pending}).Find(&flows).Error
	if err != nil {
		c.JSON(util.MakeResp(http.StatusInternalServerError, 1, err.Error()))
		return
	}
	c.JSON(util.MakeOkResp(flows))
}

func RestGet(c *gin.Context) {
	idStr := c.Param("id")
	db := DB.GetDB()
	f := new(model.Flow)
	id, _ := strconv.Atoi(idStr)
	f.ID = uint(id)
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
	c.JSON(util.MakeOkResp(f))
}

func RestGetByUser(c *gin.Context) {
	v, _ := c.Get("user")
	u := v.(model.User)
	db := DB.GetDB()
	var flows []*model.Flow
	err := db.Where(&model.Flow{ApplicantID: u.ID}).Find(&flows).Error
	if err != nil {
		c.JSON(util.MakeResp(http.StatusInternalServerError, 1, err.Error()))
		return
	}
	c.JSON(util.MakeOkResp(flows))
}
