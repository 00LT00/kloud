package user

import (
	"github.com/gin-gonic/gin"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/util"
	"net/http"
)

func RestGetInfo(c *gin.Context) {
	u, _ := c.Get("user")
	role, _ := c.Get("role")
	id := c.Query("id")
	if id == "" || u.(model.User).ID == id {
		c.JSON(util.MakeOkResp(struct {
			model.User
			Role string `json:"role"`
		}{
			u.(model.User),
			role.(string),
		}))
		return
	}
	user, err := getInfo(id)
	if err != nil {
		c.AbortWithStatusJSON(util.MakeResp(http.StatusNotFound, 1, err.Error()))
		return
	}
	c.JSON(util.MakeOkResp(*user))
}

func getInfo(userID string) (*model.User, error) {
	db := DB.GetDB()
	u := new(model.User)
	if err := db.Where(model.User{ID: userID}).First(u).Error; err != nil {
		return nil, err
	}
	u.Pass = ""
	return u, nil
}
