package go_temporary

import (
	"fmt"
)

func Main() {
	in := make(chan int, 1)
	in <- 100
	fmt.Println(<-in)
	close(in)
	fmt.Println(<-in)

	withP("1")
}

func withP(v ...string) {
	for _, t := range v {
		fmt.Println(t)
	}
}
