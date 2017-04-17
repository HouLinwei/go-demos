package main

import "fmt"
import (
	"gopkg.in/redis.v4"
	"reflect"
)
func main()  {
	fmt.Println("test redis command.")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	key := "test-set2"

	// SADD
	vids := []string{"int","float","uint","string",}
	for _ , v := range vids{
		err = client.SAdd(key, v).Err()
		if err != nil{
			fmt.Println(err)
		}
	}

	// SMEMBER
	res, err := client.SMembers(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(res))
}

