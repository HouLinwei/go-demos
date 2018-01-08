package main

import (
	"bytes"
	"fmt"
	"strings"
)

// how to use interface(io.Reader).
func main() {
	fmt.Println("----------")
	rawData := []byte{'A' - 48, 2, 3}
	reader1 := bytes.NewBuffer(rawData)
	reader2 := strings.NewReader(string(rawData))
	b1 := make([]byte, 10)
	var b2 []byte
	b2 = make([]byte, 10)
	fmt.Println(reader1.Read(b1))
	fmt.Println(reader2.Read(b2))
	fmt.Println(b1)
	fmt.Println(b2)
}
