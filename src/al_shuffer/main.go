package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//raw := []int{0,1,2,3,4,5,6,7,8,9,10}
	res := shuffler(10,19)
	fmt.Println(res)
}



func shuffler(n, needcount int )[]int{
	//
	r := []int{}
	for i := 0 ; i< n ; i++{
		r=append(r, i)
	}

	for i:= len(r)-1; i >=0; i--{
		rand.Seed(time.Now().Unix())
		j := rand.Intn(i+1)
		r[j], r[i] = r[i], r[j]
	}
	if needcount > n{
		return r
	}
	return r[0:needcount]
}
