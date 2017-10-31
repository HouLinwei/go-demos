package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	Code int
	Msg  string
}

func (e MyError) Error() string {
	e.When = time.Now()
	return fmt.Sprintf("At: %s, Code: %d, Msg: %s.", e.When, e.Code, e.Msg)
}

func Cause(c int, s string) error {
	return MyError{
		Code: c,
		Msg:  s,
	}
}

func main() {
	if err := Cause(-1, "Maybe Something Wrong."); err != nil {
		fmt.Println(err)
	}
}
