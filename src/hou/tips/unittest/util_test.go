package unittest

import (
	"testing"
)

func TestCompare(t *testing.T) {
	x := 1
	y := 2
	if Compare(x, y) {
		t.Error("wrong answear.")
	}
}
