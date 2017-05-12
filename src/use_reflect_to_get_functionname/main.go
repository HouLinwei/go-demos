package main

import (
	"fmt"
	"runtime"
)

func getFuncName() string {
	pt, _, _, _ := runtime.Caller(2)
	desc := runtime.FuncForPC(pt)
	return desc.Name()
}

func wrapper() {
	defer func() {
		name := getFuncName()
		fmt.Println(name)
	}()
}

func main() {
	wrapper()
	fmt.Println(getFuncName())
}
