package user

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"kloud/model"
	"kloud/pkg/DB"
	"kloud/pkg/redis"
	"kloud/pkg/util"
	"log"
	"net/http"
	"time"
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
	uid, _ := uuid.NewV4()
	token := uid.String()
	uJson, _ := json.Marshal(u)
	client := redis.GetRedisClient()
	err = client.Set(token, string(uJson), 24*time.Hour).Err()
	if err != nil {
		log.Println(err)
		c.JSON(util.MakeResp(http.StatusInternalServerError, 0, "unknown error"))
		return
	}
	role := getRole(u.ID)
	c.JSON(http.StatusOK, struct {
		Token string
		Role  string
	}{
		Token: token,
		Role:  role,
	})
}

func RestLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Delete("user")
	s.Delete("label")
	_ = s.Save()
	c.JSON(util.MakeOkResp("success logout"))
}
