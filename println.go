package println

import "println/logs"

type Println struct {
	Logger logs.Logger
}

const (
	PREFIX = "println"
)

func newPrintln() *Println {
	var p = &Println{}

	///var l Log

	return p
}

func Run() {
	newPrintln()
}
