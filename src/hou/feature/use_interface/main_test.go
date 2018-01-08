package use_interface

import (
	"fmt"
	"testing"
)

func TestMyWorker_GetInfo(t *testing.T) {
	my := MyWorker{}
	fmt.Println(my.GetInfo())
}
