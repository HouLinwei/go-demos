package use_interface

import "fmt"

type Worker interface {
	GetInfo() string
}

type MyWorker struct {
}

func (w *MyWorker) GetInfo() string {
	return fmt.Sprintf("[Info]:%x", w)
}
