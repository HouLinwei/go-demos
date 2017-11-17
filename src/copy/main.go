package main

import "fmt"

type P struct {
	Name string
	Age  int
	T    *T
}

type T struct {
	Name string
	Age  int
}

func (p *P) DeepCopy() *P {
	//tmp := p
	//obj := *tmp
	obj := *p
	return &obj
}

func main() {
	T := &T{"xxx", 100}
	p := P{"Houlinwei", 18, T}
	p2 := p.DeepCopy()
	p2.T.Age = 101
	fmt.Println(p.T.Age, "==", p2.T.Age)
}
