package remove_mi

import (
	"fmt"
	"strings"
)

func Remove(s string) string{
	var m = "m"
	var i = "i"
	var stack []string
	for _, v := range s {
	 t := string(v)
		if t == m{
			stack = append(stack, t)
		}else if t == i{
			if stack[len(stack)-1] == m{
				stack = stack[:len(stack)-1]
			}else{
				stack = append(stack, t)
			}
		}else{
			stack = append(stack, t)
		}
	}
	fmt.Println(stack)

	var r strings.Builder
	for _, v := range stack{
		r.WriteString(v)
	}
	return r.String()
}