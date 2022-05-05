package rest

import (
	"kloud/internal/flow"
	"kloud/internal/resource"
	"kloud/internal/user"
	"kloud/pkg/casbin"
)

func initRouter() {
	router.POST("/register", user.RestRegister)
	router.PUT("/login", user.RestLogin)
	router.PUT("/logout", user.RestLogout)

	router.Use(checkLogin)
	// 用户相关的
	{
		u := router.Group("/user")
		// 允许的标签
		u.GET("/label", user.RestLabel)
		// 获取基本信息
		u.GET("/info", user.RestGetInfo)
		// 管理员权限的增删
		a := u.Group("/admin", checkRole(casbin.Super))
		{
			a.GET("", user.RestGetAdmin)
			a.DELETE("/:id", user.RestDeleteAdmin)
			a.PATCH("/:id", user.RestAddAdmin)
		}
	}

	//资源相关
	{
		r := router.Group("/resource")
		// 创建资源
		r.POST("", checkOp("resource", casbin.Import), resource.RestCreate)
		// 获取资源
		r.GET("", resource.RestGetAll)
		r.GET("/:id", resource.RestGet)
		// 修改资源
		r.PUT("/:id", checkOp("resource", casbin.Import), resource.RestUpdate)
		// 删除资源
		r.DELETE("/:id", checkOp("resource", casbin.Import), resource.RestDelete)
	}

	//审批流
	{
		f := router.Group("/flow")
		//创建审批流
		f.POST("", flow.RestCreate)
		//获取审批详情
		f.GET("/all/:id", flow.RestGet)
		//用户自己的审批
		f.GET("/user", flow.RestGetByUser)

		//获取待审批的资源
		f.GET("/pending", checkOp("flow", casbin.Approve), flow.RestGetPending)
		//审批某个flow
		f.PUT("/id", checkOp("flow", casbin.Approve), flow.RestApprove)
	}

}
