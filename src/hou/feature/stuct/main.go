package main

import (
	"fmt"
	"time"
)

type Human struct {
	Name   string
	Birth  time.Time
	Gender string
}

func (h Human) String() string {
	var res string
	res += fmt.Sprintf("Name:\t%s\n", h.Name)
	res += fmt.Sprintf("Gender:\t%s\n", h.Gender)
	res += fmt.Sprintf("Birth:\t%s\n", h.Birth)
	return res
}

func main() {
	h := &Human{"Adam", time.Now(), "unkown"}
	fmt.Println(h)
}
