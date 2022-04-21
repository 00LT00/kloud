package rest

import "kloud/internal/user"

func initRouter() {
	u := r.Group("/user")
	u.POST("/register", user.RestRegister)
	u.PUT("/login", user.RestLogin)
	u.Use(user.InfoMiddleware())
	u.GET("/info", user.RestGetInfo)
	u.PATCH("/admin", check("admin", "add"), user.RestAddAdmin)
}

func Run(addr string) error {
	return r.Run(addr)
}
