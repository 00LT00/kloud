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

func sessionMiddleWare() gin.HandlerFunc {
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

func check(obj, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		v, _ := c.Get("user")
		u := v.(model.User)
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
