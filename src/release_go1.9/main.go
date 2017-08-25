package release_go1_9

import (
	"sync"
	"fmt"
)

func Add(x, y int) (z int) {
	z = x+y
	return
}

func ConcurrentMap(){
	m := sync.Map{}
	m.Store("1", "2")
	m.Store("2", "2")
	m.Store("3", "2")
	s, _ := m.Load("2")
	fmt.Println("Load2", s)
	m.Delete("1")
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

}