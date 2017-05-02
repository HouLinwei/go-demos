package main

import (
	"os"
	"os/signal"
	"fmt"
	"time"
)

func main()  {
	num := 2
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	in := make(chan int,num)
	outs := []chan int{}
	for i:=0;i<num;i++{
		out := make(chan int,num)
		outs = append(outs, out)
		go func(in chan int, out chan int, idx int) {
			//
			flag := false
			go func() {
				<-in
				fmt.Println("I have to terminate myself.")
				if idx == 1{
					// 这样也会导致失效。
				}else {
					out <- idx
					flag = true
				}
			}()

			for !flag{
				fmt.Println("I am running, ", idx)
				time.Sleep(time.Second*3)
			}
		}(in, out, i)
	}
	if <-c == os.Interrupt{
		fmt.Println("receive Ctrl+C.")
		close(in)
	}

	for _, i := range outs{
		idx := <-i
		if idx > -1{
			fmt.Println("end: ", idx)
		}
	}
	fmt.Println("end all.")
	return
}