package unittest

import (
	"fmt"
	"os"
	"testing"
)

func SetUp() (int, bool) {
	// do something
	return 0, true
}

func TearDown() (int, bool) {
	// tear-down
	return 0, true
}

func TestMain(m *testing.M) {
	fmt.Println("test main")
	if ret, ok := SetUp(); !ok {
		os.Exit(ret)
	}
	retCode := m.Run()
	if ret, ok := TearDown(); !ok {
		os.Exit(ret)
	}
	os.Exit(retCode)
}
