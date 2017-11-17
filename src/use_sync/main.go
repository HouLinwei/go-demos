package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	//fmt.Println()
	//a := int32(1)
	//atomic.AddInt32(&a, 1)
	//fmt.Println(a)
	//atomic.StoreInt32(&a, 100)
	//fmt.Println(a)
	//atomic.LoadInt32(&a)
	//fmt.Println(a)
}
