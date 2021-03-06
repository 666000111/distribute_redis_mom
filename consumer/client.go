package main

import "github.com/go-redis/redis"

var redisClient *redis.Client

func InitClient() (err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return nil
}
