package user

import (
	"kloud/pkg/casbin"
	"kloud/pkg/util"

	"github.com/gin-gonic/gin"
)

var (
	userLabel  = []string{"/resource", "/app"}
	adminLabel = []string{"/approval"}
	superLabel = []string{"/super"}
)

func RestLabel(c *gin.Context) {
	label, _ := c.Get("label")
	c.JSON(util.MakeOkResp(label))
}

func getLabel(id string) []string {
	e := casbin.GetEnforcer()
	roles, _ := e.GetRolesForUser(id)
	label := userLabel
	for _, role := range roles {
		tmp := make([]string, 0)
		switch role {
		case casbin.Admin:
			tmp = adminLabel
		case casbin.Super:
			tmp = superLabel
		}
		label = append(label, tmp...)
	}
	return label
}
