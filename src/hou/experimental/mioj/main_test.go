package main

import (
	"testing"
	"fmt"
)

type Case struct {
	Input string
	Output string
}

func TestSolution4(t *testing.T) {

	cases := []Case{
		Case{
			"1", "1",
		},
		Case{
			"1,2,3", "2",
		},
		Case{
			"4,5,6,7,0,1,2", "4",
		},
		Case{
			"12,13,14,5,6,7,8,9,10", "9",
		},

	}
	for _, c := range cases{
		fact := solution4(c.Input)
		if c.Output != fact{
			t.Error(fact)
		}else{
			fmt.Println("OK!")
		}
	}


}

func TestQS(t *testing.T) {
	a := []int{1,2,32,52,63,67,2,32,63,43}
	QS(a, 0, len(a)-1)
	fmt.Println(a)
}

func TestSolution2(t *testing.T){
	s := "1231231237812739878951331231231237812739878951331231231237812739878951331231231237812739878951331231231237812739878951331231231237812739870 - 89513312312312378127398789513312312312378127398789513312312312378127398789513"
    r:= "1231231237812739878951331231231237812739878951331231231237812650365639018918853110413950365639018918853110413950365639018918853110413950357"
	//s := "53121314121 - 2753538"
	//r := "53118560583"
	rr := solution2(s)
	if rr != r {
		t.Error(rr)
	}else {
		fmt.Println("OK!")
	}
}

func TestSolution3(t *testing.T){
	s := "1,2,3,4,5,6"
	r:= "6"
	rr := solution3(s)
	if rr != r {
		t.Error(rr)
	}else {
		fmt.Println("OK!")
	}
}

