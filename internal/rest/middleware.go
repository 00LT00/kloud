package rest

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"kloud/model"
	"kloud/pkg/casbin"
	"kloud/pkg/conf"
	"kloud/pkg/util"
	"log"
	"net/http"
)

func sessionMiddleware() gin.HandlerFunc {
	c := conf.GetConf()
	store, err := redis.NewStore(10, "tcp", c.Redis.Addr(), c.Redis.Pass, []byte(c.Service.Secret))
	store.Options(sessions.Options{
		MaxAge:   60 * 60 * 24,
		HttpOnly: true,
	})
	if err != nil {
		panic(err)
	}
	return sessions.Sessions("sessions", store)
}
func getUser(c *gin.Context) model.User {
	v, _ := c.Get("user")
	return v.(model.User)
}

var checkLogin = func(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("user")
	if v == nil {
		c.AbortWithStatusJSON(util.MakeResp(http.StatusUnauthorized, 0, "unauthorized"))
		return
	}
	c.Set("user", v)
	v = session.Get("label")
	if v == nil {
		c.AbortWithStatusJSON(util.MakeResp(http.StatusUnauthorized, 0, "unauthorized"))
		return
	}
	c.Set("label", v)
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
