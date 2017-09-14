package al

import (
	"testing"
	"fmt"
)

func TestShuffler(t *testing.T) {
	ar := []int{1,2,3,4,5,6,7,8,9,10}
	Shuffler(ar)
	fmt.Println(ar)
}