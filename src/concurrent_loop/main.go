package main

import (
	"fmt"
	"strings"
)

type P struct {
	Name string
	Age  int
}

func main() {
	var ps []*P

	// init
	ps = append(ps, &P{
		Name: "Hou",
		Age:  10,
	})
	ps = append(ps, &P{
		Name: "Hou",
		Age:  10,
	})
	ps = append(ps, &P{
		Name: "Hou",
		Age:  10,
	})
	ps = append(ps, &P{
		Name: "Hou",
		Age:  10,
	})
	ps = append(ps, &P{
		Name: "Hou",
		Age:  10,
	})

	finishedChan := make(chan int, len(ps))
	for idx, v := range ps {
		go process(idx, v, finishedChan)
	}
	for i := 0; i < len(ps); i++ {
		<-finishedChan
	}

	for _, v := range ps {
		fmt.Println(v)
	}

	s := " dwnidw dwdw   "
	res := strings.TrimSpace(s)
	fmt.Println(res)
	fmt.Println(len(res))
	fmt.Println(len(s))
}

func process(idx int, p *P, signal chan int) {
	p.Name = fmt.Sprintf("%s%d", p.Name, idx)
	p.Age += idx
	signal <- idx
	//if idx !=1{
	//	signal <- idx
	//}
}
