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

func RestCreate(c *gin.Context) {
	r := new(model.Resource)
	err := c.ShouldBindJSON(r)
	if err != nil {
		c.JSON(util.MakeResp(http.StatusBadRequest, 1, err.Error()))
		return
	}
	if r.Name == "" || r.Folder == "" || (r.Type != model.K8s && r.Type != model.Helm) {
		c.JSON(util.MakeResp(http.StatusBadRequest, 0, "name or folder null or type error"))
		return
	}
	err = createResource(r)
	if err != nil {
		c.JSON(util.MakeResp(http.StatusBadRequest, 1, err.Error()))
		return
	}
	c.JSON(util.MakeOkResp("create resource success"))
}

func createResource(r *model.Resource) error {
	if !filepath.IsAbs(r.Folder) {
		return errors.New("is not absolute")
	}
	ok, _ := util.PathExists(r.Folder)
	if !ok {
		return errors.New("folder not exist")
	}
	db := DB.GetDB()
	err := db.Create(r).Error
	if err != nil {
		log.Println(err.Error())
	}
	return err
}
