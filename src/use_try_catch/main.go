package main

import (
	"fmt"
)

func doWithCatchException(i int)  {
	defer func() {
		if err := recover(); err != nil{
			fmt.Println("error: ", i)
			fmt.Println(err)
		}
	}()
	var s string
	if i%2 == 0{
		s = ""
	}else {
		s = "123"
	}
	s = s[0:len(s)-1]
}

func main()  {
	for i:=0;i<10;i++{
		doWithCatchException(i)
	}
}
