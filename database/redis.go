package database

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var Cache *redis.Client

func init() {
	Cache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := Cache.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}
}
