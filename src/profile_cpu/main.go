package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {

	runtime.MemProfileRate = 1
	fileName := "/tmp/cpuprofile.p"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	go func() {
		fmt.Println(" ===== test profile cpu ===== ")
		i := 10
		for {
			i++
			fmt.Println(fab(i))
			time.Sleep(time.Second * 1)
		}
	}()

	terminate := make(chan os.Signal)
	signal.Notify(terminate, os.Interrupt)
	<-terminate
}

func fab(x int) int {
	if x == 1 || x == 2 {
		return x
	} else {
		return fab(x-1) + fab(x-2)
	}
}
