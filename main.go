package main

import (
	"log"

	"github.com/go-redis/redis"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	redis := connect()

	err := redis.HSet("user:100", "username", "fadyGamil").Err()
	if err != nil {
		log.Fatalf("error setting user data : %v", err)
	}

	res, err := redis.HGet("user:100", "username").Result()
	if err != nil {
		log.Fatalf("error getting username : %v", err)
	}
	log.Printf("username is : %v\n", res)

	// but if we used HGETALL we will get an unexpected response, which will be not an error, it will be an empty map[] in Go and an empty object in javascript
	response, err := redis.HGetAll("user:102").Result()
	if err != nil {
		log.Fatalf("error getting all fields : %v", err)
	}
	// we should check the data first if its an empty map or not
	if len(response) == 0 {
		log.Println("user:102 doesn't exists")
	} else {
		log.Printf("all fields of user:102 : %v \n", response)
	}

	// if we try to get some key that doesn't exists, redis will return nil (the expected behaviour)
	res, err = redis.HGet("user:101", "username").Result()
	if err != nil {
		log.Fatalf("error getting username : %v", err)
	}
	log.Printf("username is : %v\n", res)

}

func connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // replace <container-ip> with the actual IP
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return client

}
