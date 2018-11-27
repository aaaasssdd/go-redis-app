package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	fmt.Println(redisClient.Set("name", "jax", 0).Result())
	fmt.Println(redisClient.Get("name").String())
}
