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
	id := c.Param("id")
	db := DB.GetDB()
	var cnt int64
	db.Model(&model.App{}).Where(&model.App{AppID: id}).Count(&cnt)
	if cnt == 0 {
		c.JSON(util.MakeResp(http.StatusNotFound, 0, "app none"))
		return
	}
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
		AppID:      id,
	}

	err := db.Create(&ports).Error
	if err != nil {
		c.JSON(util.MakeResp(http.StatusInternalServerError, 1, err.Error()))
		return
	}
	c.JSON(util.MakeOkResp("create mapping success"))
}

func RestGetPortMapping(c *gin.Context) {
	id := c.Param("id")
	ports := GetPortMapping(id)
	m := make(map[int]int)
	for _, port := range ports {
		m[port.Port] = port.TargetPort
	}
	c.JSON(util.MakeOkResp(m))
}
