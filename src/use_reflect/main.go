package main

import (
	"fmt"
	"reflect"
)

type S struct {
	Name string `necessary:"1"`
	CP   string `necessary:"2"`
	Age  int    ``
}

func testReflect(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Name
		value := v.Field(i).Interface()
		tag := t.Field(i).Tag.Get("necessary")
		fmt.Println(key, value, tag)
		fmt.Println(reflect.TypeOf(key), reflect.TypeOf(value), reflect.TypeOf(tag))
		if i == 2 {
			fmt.Println(tag == "")
		}
	}
}

func main() {
	s := S{
		Name: "Jack Ma",
		CP:   "Taobao",
		Age:  18,
	}

	testReflect(s)

	//test substring
	//s1 := "1234556"
	//s1 = s1[0:len(s1)-3]
	//fmt.Println(s1)

}
