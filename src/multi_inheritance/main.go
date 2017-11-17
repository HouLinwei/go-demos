package main

import "fmt"

type A interface {
	F()
}

type B struct {
	A
}

func (b B) F() {
	fmt.Println("F() of B")
}

type C struct {
	A
}

func (c C) F() {
	fmt.Println("F() of C")
}

type D struct {
	B
	C
}

func main() {
	d := D{}

	// Error! "ambiguous selector d.F".
	// d.F()

	// Allowed!
	d.B.F()
	d.C.F()
	d.B.A = d.C
	d.B.A.F()
}
