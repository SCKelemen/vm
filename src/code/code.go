package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Instructions []byte

func (ins Instructions) String() string {
	var out bytes.Buffer

	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ErrorL %s\n", err)
			continue
		}

		operands, read := ReadOperands(def, ins[i+1:])

		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))

		i += 1 + read
	}

	return out.String()

}

func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)

	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand count does not match definition. Expected: %d Actual: %d\n", operandCount, len(operands))
	}

	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	default:
		return fmt.Sprintf("ERROR: unhandled opernadCount for %s\n", def.Name)
	}
}

type OpCode uint8

const (
	NOP OpCode = iota

	CONST

	ADD
	SUB
	MUL
	QUO

	TRUE
	FALSE

	EQL
	NEQL

	LT
	LTE

	GT
	GTE
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[OpCode]*Definition{
	NOP: {"NOP", []int{}},

	CONST: {"CONST", []int{2}},

	ADD: {"ADD", []int{}},
	SUB: {"SUB", []int{}},
	MUL: {"MUL", []int{}},
	QUO: {"QUO", []int{}},

	TRUE:  {"TRUE", []int{}},
	FALSE: {"FALSE", []int{}},

	EQL:  {"EQL", []int{}},
	NEQL: {"NEQL", []int{}},

	GT:  {"GT", []int{}},
	GTE: {"GTE", []int{}},

	LT:  {"LT", []int{}},
	LTE: {"LTE", []int{}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[OpCode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

func Make(op OpCode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLength := 1
	for _, w := range def.OperandWidths {
		instructionLength += w
	}

	instruction := make([]byte, instructionLength)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}
	return instruction
}

func ReadOperands(def *Definition, ins Instructions) (operands []int, offset int) {
	operands = make([]int, len(def.OperandWidths))
	offset = 0

	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}

		offset += width
	}
	return
}

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}
