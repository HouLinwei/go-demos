package release_go1_9

import (
	"fmt"
	"sync"
	"time"
)

func Add(x, y int) (z int) {
	z = x + y
	return
}

func ConcurrentMap() {
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

func XTime() {
	t := time.Now()
	println(t.String())
	println("======")

	nt := t.Add(time.Hour * 1)
	println(t.String())
	println(nt.String())
	println("======")

	n2t := t.AddDate(1, 1, 1)
	println(t.String())
	println(n2t.String())
	println("======")

	n3t := t.Round(time.Hour * 1)
	println(t.String())
	println(n3t.String())
	println("======")

	n4t := t.After(time.Time{})
	println(t.String())
	println(n4t)
	println("======")
}
