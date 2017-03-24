package main

import "fmt"

const (
	// enum test
	read = iota
	write
	delete
)

func main()  {
	fmt.Print(read)
	fmt.Print(write)
	fmt.Print(delete)
}
