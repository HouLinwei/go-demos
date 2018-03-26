package print_matrix

import "testing"

func TestPrint(t *testing.T) {
	Print()
}

func TestPrint2(t *testing.T) {
	var m [4][4]int
	m = [4][4]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
		{13,14,15,16},
	}
	Print2(m, 4)
}