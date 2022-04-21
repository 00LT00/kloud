package rest

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"kloud/model"
)

var r *gin.Engine

func init() {
	gob.Register(model.User{})
	r = gin.Default()
	r.Use(sessionMiddleWare())
	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		_ = session.Save()
		c.JSON(200, gin.H{"count": count})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	initRouter()
}
