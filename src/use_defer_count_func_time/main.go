package main

import (
	"time"
	"fmt"
)

func doSth(i time.Duration){
	time.Sleep(time.Second*i)
}

func main()  {
	t1 := time.Now().Unix()
	defer func(start int64) {
		now := time.Now().Unix()
		fmt.Println("time delta:", now - start)
	}(t1)
	doSth(1)
}

