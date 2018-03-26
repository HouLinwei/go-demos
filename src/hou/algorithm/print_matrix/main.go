package print_matrix

import "fmt"

func Print(){
	var m [4][4]int
	m = [4][4]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
		{13,14,15,16},
	}
	fmt.Println(m)

	dim := len(m[0])
	//dim := 4
	depth := dim - 2
	num := dim - 1

	fmt.Printf("dim: %d, depth:%d, num:%d\n", dim, depth, num)
	for i :=0 ; i< depth; i++{
		for j :=0 ; j < dim; j++{
			for k :=0 ; k < num; k++{
				fmt.Printf("%d ", m[i][k])
			}
			fmt.Println()
			// rotate it
		}
	}

}

func rotate(raw [4][4]int, d, r, c int)[][]int{
	var ret [][]int
	return ret
}

func Print2(m [4][4]int, dim int){
	t := 0
	for dim >0 {
		fmt.Println("dim: ", dim)
		// print out margin
		for i:=0 ;i<dim;i++ {
			fmt.Printf("%d ", m[0+t][i+t])
		}
		fmt.Println()
		for i:=1; i<dim; i++ {
			fmt.Printf("%d ", m[i+t][dim])
		}
		fmt.Println()
		for i:=1; i<dim; i++ {
			fmt.Printf("%d ", m[dim-1][dim-1-i])
		}
		fmt.Println()
		for i:=2; i<dim; i++ {
			fmt.Printf("%d ", m[dim-i][0+t])
		}
		fmt.Println()
		//
		dim = dim -2
		t += 1
	}
	//if dim > 0{
	//	for i:=0 ;i<dim;i++ {
	//		fmt.Printf("%d ", m[0][i])
	//	}
	//	fmt.Println()
	//	for i:=1; i<dim; i++ {
	//		fmt.Printf("%d ", m[i][dim-1])
	//	}
	//	fmt.Println()
	//	for i:=1; i<dim; i++ {
	//		fmt.Printf("%d ", m[dim-1][dim-1-i])
	//	}
	//	fmt.Println()
	//	for i:=2; i<dim; i++ {
	//		fmt.Printf("%d ", m[dim-i][0])
	//	}
	//	fmt.Println()
	//	// update m
	//	// Print2(m, dim-2)
	//}
	//if dim == 1 {
	//	fmt.Println(m[0][0])
	//}
}