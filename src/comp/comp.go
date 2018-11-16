package comp

import (
	"github.com/sckelemen/vm/src/code"
	"github.com/sckelemen/vm/src/obj"
)

type Compilation struct {
	instructions	code.instructions
	constants		[]obj.Object
}

func New() *Compilation {
	return &Compilation{
		instructions: code.Instructions{},
		constants:	[]obj.Object{}
	}
}

func (c *Compilation) Compile(node ast.Node) error {

}

type Bytecode struct {
	Instructions	code.Instructions
	Constants		[]obj.Object
}