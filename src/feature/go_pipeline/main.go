package go_pipeline

import (
	"fmt"
	"sync"
	"time"
)

// 一个简单的Fan-Out的pipeline
func RunPipeLine() {
	in := make(chan int, 10)
	go func() {
		i := 0
		for {
			in <- i
			time.Sleep(time.Second * 1)
			i++
		}
	}()

	//for n := range in {
	//	fmt.Println("In: ", n)
	//}

	done := make(chan int, 1)
	out := Do(done, in)
	go func() {
		time.Sleep(time.Second * 5)
		//done <- 1
		close(done)
	}()
	for n := range out {
		fmt.Println("Out: ", n)
	}

	fmt.Println("Finished.")
}

func Do(done <-chan int, in <-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range in {
			select {
			// 搞清楚这儿的语义
			case out <- n:
				fmt.Println("Recv: ", n)
			case t := <-done:
				fmt.Println("DONE", t)
				return
			}
		}
	}()
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
