package remove_mi

import (
	"testing"
	"fmt"
)

func TestRemove(t *testing.T) {
	var s = "mabcdmmimiefgmmiihijmidmi"
	var target = "mabcdmefghijd"
	if Remove(s) != target{
		t.Error("wrong!")
	}else{
		fmt.Println("OK!")
	}
}
