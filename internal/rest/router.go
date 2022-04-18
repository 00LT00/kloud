package rest

func initRouter() {
	user := r.Group("/user")
	user.POST("")

}

func Run(addr string) error {
	return r.Run(addr)
}
