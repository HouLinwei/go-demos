package main

import (
	"fmt"
	"strconv"
	"encoding/json"
	"time"
	"sync"
)


type MyIn interface {
	GetName()string
}

type MyIn2 interface {
	GetName2()string
}

type myInImpl struct {
	Name string
}

func (m myInImpl)GetName()string{
	return m.Name+"1"
}

func (m *myInImpl)GetName2()string{
	return m.Name+"2"
}


func Run1(ps []MyIn){
	for _, v := range ps{
		fmt.Println(v.GetName())
	}
}

func Run2(ps []MyIn2){
	for _, v := range ps{
		fmt.Println(v.GetName2())
	}
}


func main0()  {
	fmt.Println("HELLO JUDY")

	ps1 := []MyIn{
		myInImpl{"n1",},
		myInImpl{"n2",},
	}

	Run1(ps1)

	fmt.Println("HELLO JUDY")
	//c := redis.NewClient(&redis.Options{
	//	Addr:"127.0.0.1:6379",
	//})
	//ok, er := c.Exists("test").Result()
	//fmt.Println(ok)
	//fmt.Println(er)


	//var ret *S
	//
	//f := func() {
	//	ret, ok := tx()
	//	fmt.Println("In func: ", *ret, ok)
	//}
	//
	//f()
	//
	//fmt.Println(ret)


	i64, err := strconv.ParseInt("100",2,0)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(i64)


	m := make(map[string]int,1)
	m[""] = 1

	fmt.Println(m)

	fmt.Println("=======")

	as := `
		{"id":5627075427}
	`
	var ret interface{}

	json.Unmarshal([]byte(as), &ret)
	fmt.Println(ret)


	ass := 5.627075427e+09
	ns := strconv.FormatFloat(ass, 'g',20, 64)
	fmt.Printf("%s\n",ns)
	fmt.Printf("%x\n",ass)
	fmt.Printf("%X\n",ass)

}


func tx()(*S, bool){
	return &S{1}, true
}

type S struct {
	A int
}

func main1(){
	defer fmt.Println("after main")

	go func() {
		defer func() {
			if err := recover(); err != nil{
				fmt.Println(err)
				panic("in recover")
			}
		}()

		panic("in go routine")
		fmt.Println("after panic")
	}()

	time.Sleep(time.Second*3)
	fmt.Println("in main")
}

func main(){
	var l sync.Mutex

	go func() {
		l.Lock()
		fmt.Println("locking")
		time.Sleep(time.Second*5)
		l.Unlock()
	}()

	time.Sleep(time.Millisecond*500)
	l.Lock()
	fmt.Println("xxx")
	l.Unlock()
	fmt.Println("end")

	fmt.Println(TodayDateStr())

}

func TodayDateStr() string {
	return DateStr(time.Now())
}

func DateStr(n time.Time) string {
	y, m, d := n.Date()
	//fmt.Println(y,m, d)
	return fmt.Sprintf("%d-%d-%d", y, m, d)
}
