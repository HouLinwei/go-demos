package main

import "fmt"

func Print() {

}

func main() {
	vids := []string{"1", "2"}
	fmt.Println(vids)
	tmpMap := make(map[string]string)
	tmpMap["a"] = "av"
	tmpMap["b"] = "bv"
	fmt.Println(len(tmpMap) == 2)
	for k, v := range tmpMap {
		fmt.Println(k, v)
	}
}
