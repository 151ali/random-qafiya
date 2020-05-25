package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

// ConnectRedis :
func ConnectRedis() (*redis.Client, error) {
	log.Println("Connecting to Redis ... ")
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Println("Fail to connect to Redis :(")
		return nil, err
	}
	log.Println("You are Connected to Redis :) ")
	fmt.Println(">>", RedisClient)
	return RedisClient, nil
}
