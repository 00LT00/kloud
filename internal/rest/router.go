package rest

import (
	"kloud/internal/user"
	"kloud/pkg/casbin"
)

func initRouter() {
	r.POST("/register", user.RestRegister)
	r.PUT("/login", user.RestLogin)
	r.PUT("/logout", user.RestLogout)
	// 用户相关的
	u := r.Group("/user", user.LoadInfoMiddleware())
	{
		// 允许的标签
		u.GET("/label", user.RestLabel)
		// 获取基本信息
		u.GET("/info", user.RestGetInfo)
		// 管理员权限的增删
		a := u.Group("/admin", checkRole(casbin.Super))
		{
			a.GET("/", user.RestGetAdmin)
			a.DELETE("/:id", user.RestDeleteAdmin)
			a.PATCH("/:id", user.RestAddAdmin)
		}
	}
}
