package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/util"
	"log"
	"net/http"
)

func RestLogin(c *gin.Context) {
	db := DB.GetDB()
	type req struct {
		Username, Password string
	}
	r := new(req)
	_ = c.ShouldBindJSON(r)
	username := r.Username
	password := r.Password
	if username == "" || password == "" {
		c.JSON(util.MakeResp(http.StatusBadRequest, 0, "username or password null"))
		return
	}
	u := new(model.User)
	u.Name = username
	err := db.Where(u).First(u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(util.MakeResp(http.StatusNotFound, 0, "user none"))
			return
		}
		log.Println(err)
		c.JSON(util.MakeResp(http.StatusInternalServerError, 0, "unknown error"))
		return
	}
	if !util.PasswordVerify(password, u.Pass) {
		c.JSON(util.MakeResp(http.StatusForbidden, 0, "password error"))
		return
	}
	u.Pass = ""
	s := sessions.Default(c)
	s.Set("user", u)
	label := getLabel(u.ID)
	s.Set("label", label)
	_ = s.Save()
	c.JSON(http.StatusOK, struct {
		User  *model.User
		Label []string
	}{
		User:  u,
		Label: label,
	})
}

func RestLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Delete("user")
	s.Delete("label")
	_ = s.Save()
	c.JSON(util.MakeOkResp("success logout"))
}
