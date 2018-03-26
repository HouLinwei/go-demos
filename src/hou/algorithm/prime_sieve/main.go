package main

// A concurrent prime sieve

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	//fmt.Println(ch1)
	go func(in chan int) {
		i := 1
		for {
			i++
			in <- i
		}
	}(ch1)
	for i := 0; i < 10; i++ {
		prime := <-ch1
		fmt.Println(prime)
		ch2 := make(chan int)
		//fmt.Println(ch2)
		//fmt.Println("addr ch1:", ch1)
		go func(in, out chan int, prime int) {
			for {
				i := <-in
				if i%prime != 0 {
					out <- i
				}
			}
		}(ch1, ch2, prime)
		ch1 = ch2
		//fmt.Println("addr ch1:", ch1)
		//fmt.Println("addr ch2:", ch2)
	}
}
