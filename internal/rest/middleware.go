package rest

import (
	"encoding/json"
	"kloud/model"
	"kloud/pkg/casbin"
	"kloud/pkg/redis"
	"kloud/pkg/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) model.User {
	v, _ := c.Get("user")
	return v.(model.User)
}

var checkLogin = func(c *gin.Context) {
	token := c.GetHeader("Authorization")
	redisClient := redis.GetRedisClient()
	uStr, _ := redisClient.Get(token).Result()
	if uStr == "" {
		c.AbortWithStatusJSON(util.MakeResp(http.StatusUnauthorized, 0, "unauthorized"))
		return
	}
	v := new(model.User)
	_ = json.Unmarshal([]byte(uStr), v)
	if v == nil {
		c.AbortWithStatusJSON(util.MakeResp(http.StatusUnauthorized, 0, "unauthorized"))
		return
	}
	c.Set("user", *v)
}

func checkOp(obj, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := getUser(c)
		e := casbin.GetEnforcer()
		ok, err := e.Enforce(u.ID, obj, act)
		if !ok || err != nil {
			if err != nil {
				log.Println(err.Error())
			}
			c.AbortWithStatusJSON(util.MakeResp(http.StatusForbidden, 0, "forbidden"))
			return
		}
	}
}

func checkRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := getUser(c)
		e := casbin.GetEnforcer()
		for _, r := range roles {
			ok, _ := e.HasRoleForUser(u.ID, r)
			if ok {
				return
			}
		}
		c.AbortWithStatusJSON(util.MakeResp(http.StatusForbidden, 0, "forbidden"))
		return
	}
}
