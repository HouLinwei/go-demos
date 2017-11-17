package main

import (
	"fmt"
	"runtime"
	"time"
)

func getFuncName() string {
	pt, _, _, _ := runtime.Caller(2)
	desc := runtime.FuncForPC(pt)
	return desc.Name()
}

func count(start int64) {
	time.Sleep(1 * time.Second)
	now := time.Now().UnixNano()
	fmt.Println(float64(now-start) / 1e9)
}

func wrapper() {
	defer count(time.Now().UnixNano())
}

func main() {
	wrapper()
	fmt.Println(getFuncName())
}
