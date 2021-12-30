package main

import (
	"github.com/go-redis/redis/v7"
)

var redisdb *redis.Client

func GetRedis() *redis.Client {
	if redisdb == nil {
		redisdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	}
	return redisdb
}
