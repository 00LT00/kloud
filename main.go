package main

import (
	"kloud/internal/rest"
	"kloud/pkg/DB"
	"kloud/pkg/casbin"
	"kloud/pkg/conf"
	"kloud/pkg/redis"
	"log"
	"time"
)

func main() {
	log.Println("hello kloud")
	conf.ShowAllConfig()
	DB.Ping()
	e := casbin.GetEnforcer()
	if e == nil {
		log.Println("nil")
	}
	log.Println(e.GetPolicy())
	redis.GetRedisClient().Set("kloud", time.Now().Format(time.RFC3339), 0)
	log.Println(redis.GetRedisClient().Get("kloud").Time())
	rest.Run(conf.GetConf().Service.Addr())
}
