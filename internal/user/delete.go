package user

import (
	"github.com/gin-gonic/gin"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/casbin"
	"kloud/pkg/util"
	"net/http"
)

func RestDeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(util.MakeResp(http.StatusBadRequest, 0, "id none"))
		return
	}

	if err := deleteUser(id); err != nil {
		c.JSON(util.MakeOkResp("success"))
	} else {
		c.JSON(util.MakeResp(http.StatusOK, 1, err.Error()))
	}
}

func deleteUser(ID string) error {
	e := casbin.GetEnforcer()
	db := DB.GetDB()
	tx := db.Begin()
	err := tx.Delete(&model.User{}, ID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = e.DeleteUser(ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
