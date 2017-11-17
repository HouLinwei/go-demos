package main

import (
	"fmt"
	"gopkg.in/redis.v4"
	"sync/atomic"
	"time"
)

const (
	host = "10.108.35.17"
	port = 8991
)

func main() {
	fmt.Println("=========")

	date := time.Now().Format("2006-01-02-15")

	fmt.Println(date)

	fmt.Println("=========")
	var myAtomic atomic.Value
	myAtomic.Store("xxx")
	x := myAtomic.Load()
	fmt.Println(x)

	y, z := f1, f2
	fmt.Println(y, z)
	y("a")
	z("b")

	fmt.Println("=========")
	now := time.Now().Add(-time.Hour * 24)
	r := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
	fmt.Println(r)

	fmt.Println("=========")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	k := "xxxx"
	keys, _ := client.HKeys(k).Result()
	//client.HSet(k, "v11","1")
	//client.HSet(k, "v12","1")
	//client.HSet(k, "v13","1")
	//client.HSet(k, "v14","1")
	fmt.Println(keys)
	res := client.HDel(k, keys...)
	fmt.Println(res)

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)
}

func f1(s string) {
	fmt.Println("f1")
}

func f2(s string) {
	fmt.Println("f2")
}
