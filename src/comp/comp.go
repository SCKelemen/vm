package comp

import "code"

type Compilation struct {
	instructions code.instructions
}

func New() *Compilation {
	return &Compilation{
		instructions: code.Instructions{},
		constants:	[]object.Object{}
	}
}