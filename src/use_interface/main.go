package use_interface

type Worker interface {
	GetInfo()string
}

func (w *Worker)GetInfo()string{
	return ""
}