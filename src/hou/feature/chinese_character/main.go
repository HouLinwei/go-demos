package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// https://golang.org/pkg/unicode/
func main() {
	s := "这是中文。"
	// Get the length of s as unicode.
	fmt.Println("length is: ", utf8.RuneCountInString(s))
	for _, value := range s {
		// Whether a Chinese character.
		fmt.Println(unicode.Is(unicode.Scripts["Han"], value))
	}
}
