package x_update

import (
	"fmt"
	"sync"
	"time"
)

var t int

type MyVar struct {
	l    sync.RWMutex
	Data int
}

type MyVar2 struct {
	l    sync.RWMutex
	Data map[string]string
}

func Run() {
	go func() {
		ticker := time.NewTicker(time.Second * 1)
		defer ticker.Stop()
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()
	m := make(map[string]string)
	m = map[string]string{
		"1": "1",
	}
	fmt.Println(m)
	UpdateM(m)
	fmt.Println(m)
	v2 := MyVar2{}
	v2.Data = m
	fmt.Println(v2)
	fmt.Println("End")
}

func UpdateM(m map[string]string) {
	m["2"] = "2"
}

func RunWithLock() {
	v := MyVar{
		Data: 1,
	}
	for i := 0; i < 100000; i++ {
		printVWithLock(v)
	}
}

func printVWithLock(v MyVar) {
	v.l.RLock()
	defer v.l.RUnlock()
	fmt.Println(v.Data)
}

func RunWithoutLock() {
	v := MyVar{
		Data: 1,
	}
	for i := 0; i < 100000; i++ {
		printVWithoutLock(v)
	}
}

func printVWithoutLock(v MyVar) {
	fmt.Println(v.Data)
}
