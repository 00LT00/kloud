package port

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/k8s"
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
	port := &model.PortMapping{
		Port:       req.Port,
		TargetPort: req.TargetPort,
		AppID:      id,
	}
	err := k8s.ForwardPort("pods", port.AppID, port.Port, port.TargetPort)
	if err != nil {
		c.JSON(util.MakeResp(http.StatusBadRequest, 1, err.Error()))
		return
	}
	err = db.Create(&port).Error
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

func RestDeletePortMapping(c *gin.Context) {
	id := c.Param("id")
	db := DB.GetDB()
	var cnt int64
	db.Model(&model.App{}).Where(&model.App{AppID: id}).Count(&cnt)
	if cnt == 0 {
		c.JSON(util.MakeResp(http.StatusNotFound, 0, "app none"))
		return
	}
	port := c.Param("port")
	err := deletePortMapping(id, util.Str2Int(port))
	if err != nil {
		c.JSON(util.MakeResp(http.StatusInternalServerError, 1, err.Error()))
		return
	}
	c.JSON(util.MakeOkResp("delete mapping success"))
}

func deletePortMapping(id string, port int) (err error) {
	db := DB.GetDB()
	var cnt int64
	db.Model(&model.PortMapping{}).Where(&model.PortMapping{AppID: id, Port: port}).Count(&cnt)
	if cnt == 0 {
		return errors.New("port mapping none")
	}
	k8s.StopPort(port)
	err = db.Delete(&model.PortMapping{AppID: id, Port: port}).Error
	return
}
