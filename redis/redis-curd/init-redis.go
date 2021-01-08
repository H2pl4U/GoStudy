package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

//声明一个全局的rdb变量
var rdb *redis.Client

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  //no pwd set
		DB:       0,   //use default DB
		PoolSize: 100, //连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = rdb.Ping(ctx).Result()
	return err
}

//V8Example V8新版本升级
func V8Example() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		return
	}

	err := rdb.Set(ctx, "key234", "value234", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "key234").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key233").Result()
	if err == redis.Nil {
		fmt.Println("key233 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key233", val2)
	}
}

func main() {
	initClient()
	V8Example()
}
