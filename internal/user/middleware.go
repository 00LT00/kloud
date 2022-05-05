package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"kloud/pkg/util"
	"net/http"
)

func LoadInfoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}
