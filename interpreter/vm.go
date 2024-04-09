package interpreter

import (
	"my-language/ast"
)

// VM represents the virtual machine for the language.
type VM struct {
	instructions []Instruction
	stack        []interface{}
	globals      map[string]interface{}
}

// Instruction represents a virtual machine instruction.
type Instruction struct {
	Opcode   OpCode
	Operand1 interface{}
	Operand2 interface{}
}

// OpCode represents the opcode of an instruction.
type OpCode int

const (
	OpPush OpCode = iota
	OpPop
	OpAdd
	OpSub
	OpMul
	OpDiv
	// Add more opcodes as needed
)

// Run executes the program on the VM.
func (vm *VM) Run() error {
	// Implement the fetch-decode-execute cycle of the VM
	return nil
}

// Load loads the program into the VM.
func (vm *VM) Load(node ast.Node) error {
	// Implement the logic to translate the AST into VM instructions
	return nil
}
