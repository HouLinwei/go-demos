package main

import "fmt"

func main() {

	MyFunc := func(p1, p2 string) func(string) {
		cnt := 0
		return func(p string) {
			cnt ++
			fmt.Println(p1, p2, p, cnt)
		}
	}

	f1 := MyFunc("a", "b")
	f1("c")
	f1("d")
	f1("d")

	f2 := MyFunc("h", "i")
	f2("c")
	f2("d")
	f2("e")
}
