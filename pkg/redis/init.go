package redis

import (
	"context"
	goredis "github.com/go-redis/redis/v8"
	"kloud/pkg/conf"
	"sync"
	"time"
)

type RedisClient struct {
	ctx context.Context
	rdb *goredis.Client
}

var r *RedisClient

var once sync.Once

func initRedisClient() {
	r = new(RedisClient)
	c := conf.GetConf()
	r.rdb = goredis.NewClient(&goredis.Options{
		Addr:     c.Redis.Addr(),
		Password: c.Redis.Pass, // no password set
		DB:       0,            // use default DB
	})
	r.ctx = context.Background()
}

func GetRedisClient() *RedisClient {
	once.Do(func() {
		initRedisClient()
	})
	return r
}

func (r RedisClient) Get(key string) *goredis.StringCmd {
	return r.rdb.Get(r.ctx, key)
}

func (r RedisClient) Set(key string, value interface{}, expiration time.Duration) *goredis.StatusCmd {
	return r.rdb.Set(r.ctx, key, value, expiration)
}
