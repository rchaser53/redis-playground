package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := createRedisClient()
	ExampleClient(client)
}

func createRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return client
}

// ExampleClient is nya-n
func ExampleClient(client *redis.Client) {
	err := client.Set("key", "nyan", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
