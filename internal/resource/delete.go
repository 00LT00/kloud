package resource

import (
	"github.com/gin-gonic/gin"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/util"
	"log"
	"net/http"
)

func RestDelete(c *gin.Context) {
	id := c.Param("id")

	err := deleteResource(id)
	if err != nil {
		c.JSON(util.MakeResp(http.StatusBadRequest, 1, err.Error()))
		return
	}
	c.JSON(util.MakeOkResp("delete resource success"))
}

func deleteResource(id string) error {
	db := DB.GetDB()
	err := db.Delete(&model.Resource{ResourceID: id}).Error
	if err != nil {
		log.Println(err.Error())
	}
	return err
}
