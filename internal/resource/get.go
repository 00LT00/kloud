package resource

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/util"
	"log"
	"net/http"
)

func RestGetAll(c *gin.Context) {
	db := DB.GetDB()
	var rs []*model.Resource
	err := db.Find(&rs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(util.MakeResp(http.StatusInternalServerError, 1, err.Error()))
		return
	}
	c.JSON(util.MakeOkResp(rs))
}

func RestGet(c *gin.Context) {
	id := c.Param("id")
	db := DB.GetDB()
	r := new(model.Resource)
	r.ResourceID = id
	err := db.Where(r).First(r).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(util.MakeResp(http.StatusNotFound, 0, "resource none"))
			return
		}
		log.Println(err)
		c.JSON(util.MakeResp(http.StatusInternalServerError, 0, "unknown error"))
		return
	}
	c.JSON(util.MakeOkResp(struct {
		model.Resource
		Config string `json:"config"`
	}{
		*r,
		`{"name":"demo","cpu":"20","memory": "10"}`,
	}))
}
