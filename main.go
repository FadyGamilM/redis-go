package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // replace <container-ip> with the actual IP
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
