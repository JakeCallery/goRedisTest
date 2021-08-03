package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	fmt.Println("start")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()

	if err != nil {
		log.Fatalf("[ERROR]: %v", err)
	}

	fmt.Printf("Pong: %v\n", pong)
}
