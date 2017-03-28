package main

import (
	"fmt"
	"time"
)

func CallWithVariableParameters(names ...string) {
	for index, value := range names {
		fmt.Printf("ID: %d, Value: %s.\n", index, value)
	}
	fmt.Println()
}

type Parameter struct {
	Name   string
	Birth  time.Time
	Gender int
	// ...
}

func CallWithOptionalParameters(param Parameter) {
	fmt.Println(param.Name, param.Birth, param.Gender)
}

func main() {
	// Variable Parameters
	CallWithVariableParameters("yebidaxiong")
	names := []string{
		"yebidaxiong", "duolaameng", "jingxiang",
	}
	CallWithVariableParameters(names...)

	// Optional Parameters. Not cool.
	p1 := Parameter{"BigX", time.Now(), 1}
	CallWithOptionalParameters(p1)

	p2 := Parameter{
		Birth: time.Now(),
	}
	CallWithOptionalParameters(p2)
}
