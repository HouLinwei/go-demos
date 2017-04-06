package main

import (
	"fmt"
)

func main() {
	s := "123abc!!!你好！"
	fmt.Println("==========")
	for _, v := range s {
		fmt.Printf("%c \t", v)
	}
	fmt.Println()
	for _, v := range s {
		fmt.Printf("%d \t", v)
	}
	fmt.Println()
	fmt.Println("==========")

	b := []byte(s)
	for _, v := range b {
		fmt.Printf("%c \t", v)
	}
	fmt.Println()
	for _, v := range b {
		fmt.Printf("%d \t", v)
	}
	fmt.Println()
	fmt.Println("==========")

	r := []rune(s)
	for _, v := range r {
		fmt.Printf("%c \t", v)
	}
	fmt.Println()
	for _, v := range r {
		fmt.Printf("%d \t", v)
	}
	fmt.Println()
	fmt.Println("==========")
}
