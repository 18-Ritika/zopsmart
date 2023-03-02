package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	res, err := client.Ping().Result()
	fmt.Println(res, err)

	er := client.Set("id123", "Ritika", 0).Err()
	if err != nil {
		fmt.Println(er)
	}

	value, err := client.Get("id123").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)

}
