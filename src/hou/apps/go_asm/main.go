package main

import "fmt"

func neg(x uint64) int64

func main(){
	fmt.Println("Go ASM!")
	s := neg(1)
	fmt.Println(s)
}
