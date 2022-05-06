package resource

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/util"
	"log"
	"net/http"
	"path/filepath"
)

func RestUpdate(c *gin.Context) {
	id := c.Param("id")
	r := new(model.Resource)
	err := c.ShouldBindJSON(r)
	if err != nil {
		c.JSON(util.MakeResp(http.StatusBadRequest, 1, err.Error()))
		return
	}
	if r.Name == "" || r.Folder == "" {
		c.JSON(util.MakeResp(http.StatusBadRequest, 0, "name or folder null"))
		return
	}
	r.ResourceID = id
	err = updateResource(r)
	if err != nil {
		c.JSON(util.MakeResp(http.StatusBadRequest, 1, err.Error()))
		return
	}
	c.JSON(util.MakeOkResp("update resource success"))
}

func updateResource(r *model.Resource) error {
	if !filepath.IsAbs(r.Folder) {
		return errors.New("is not absolute")
	}
	ok, _ := util.PathExists(r.Folder)
	if !ok {
		return errors.New("folder not exist")
	}
	db := DB.GetDB()
	err := db.Where(&model.Resource{ResourceID: r.ResourceID}).Save(r).Error
	if err != nil {
		log.Println(err.Error())
	}
	return err
}
