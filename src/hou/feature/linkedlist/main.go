package main

import (
	"container/list"
	"fmt"
)

func main() {
	linkedList := list.New()
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	linkedList.PushFront(0)
	linkedList.InsertAfter(2.2, linkedList.Front().Next().Next())
	fmt.Println("len:", linkedList.Len())
	for e := linkedList.Front(); e != nil; e = e.Next() {
		fmt.Print("{", e.Value, "}") //输出list的值
		if e.Next() != nil {
			fmt.Print("->")
		}
	}
}
