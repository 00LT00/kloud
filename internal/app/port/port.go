package port

import (
	"github.com/gin-gonic/gin"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/util"
	"net/http"
)

func GetPortMapping(id string) []model.PortMapping {
	db := DB.GetDB()
	var portMappings []model.PortMapping
	db.Where(&model.PortMapping{AppID: id}).Find(&portMappings)
	return portMappings
}

func RestCreatePortMapping(c *gin.Context) {
	var req struct {
		Port, TargetPort int
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(util.MakeResp(http.StatusBadRequest, 1, err.Error()))
		return
	}
	ports := &model.PortMapping{
		Port:       req.Port,
		TargetPort: req.TargetPort,
	}
	db := DB.GetDB()
	err := db.Create(&ports).Error
	if err != nil {
		c.JSON(util.MakeResp(http.StatusInternalServerError, 1, err.Error()))
		return
	}
	c.JSON(util.MakeOkResp("create mapping success"))
}
