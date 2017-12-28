package x_update

import (
	"testing"
	"fmt"
)

func TestRun(t *testing.T) {
	Run()
	fmt.Println("End")
}

func BenchmarkRunWithLock(b *testing.B) {
	b.N = 100
	for i:=0; i< b.N;i++{
		RunWithLock()
	}
}

func BenchmarkRunWithoutLock(b *testing.B) {
	b.N = 100
	for i:=0; i< b.N;i++{
		RunWithoutLock()
	}
}