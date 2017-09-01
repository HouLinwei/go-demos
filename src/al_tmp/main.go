package main

import (
	"os/exec"
	"fmt"
)

func main()  {
	println("Justice.")
	path, err := exec.LookPath("go")
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(path)
	}
}
