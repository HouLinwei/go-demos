package release_go1_9

import (
	"testing"
	"fmt"
)

func help(t *testing.T){
	t.Helper()
	fmt.Println("put help info here")
}

func TestAdd(t *testing.T) {
	help(t)
}

func TestConcurrentMap(t *testing.T) {
	ConcurrentMap()
}

func TestXTime(t *testing.T) {
	XTime()
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
		Add(1,2)
	}
}