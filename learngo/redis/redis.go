package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "121.40.151.71:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})

	return
}

func Close() {
	_ = rdb.Close()
}

func main() {
	err := Init()
	if err != nil {
		log.Fatal(err)
	}
	defer rdb.Close()

	var ctx = context.Background()

	err = rdb.Set(ctx, "rey", "111", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "rey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("rey", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
