package cache

import (
	"Todo/conf"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func LinkRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPassword,
		DB:       conf.RedisDB,
	})
	RedisClient = rdb
}
