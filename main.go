package main

import (
	"fmt"
	"kloud/pkg/casbin"
	"kloud/pkg/conf"
	"kloud/pkg/database"
	"kloud/pkg/redis"
	"time"
)

func main() {
	fmt.Println("hello kloud")
	conf.ShowAllConfig()
	database.Ping()
	e := casbin.GetEnforcer()
	if e == nil {
		fmt.Println("nil")
	}
	redis.GetRedisClient().Set("kloud", time.Now().Format(time.RFC3339), 0)
	fmt.Println(redis.GetRedisClient().Get("kloud").Time())
}
