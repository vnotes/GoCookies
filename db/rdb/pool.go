package rdb

import "github.com/go-redis/redis/v8"

var Pool *redis.Client

func init() {
	Pool = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
