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