package infrastructure

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func RedisClient() *redis.Client {
	return rdb
}

func Ctx() context.Context {
	return ctx
}
