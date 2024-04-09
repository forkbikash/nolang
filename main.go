package main

import (
	"my-language/internal/utils"
	"my-language/interpreter"
	"my-language/parser"
)

func main() {
	// Parse the input file
	filename := "internal/testdata/example.my"
	source, err := utils.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Parse the source code
	parser := &parser.Parser{}
	root, err := parser.Parse(string(source))
	if err != nil {
		panic(err)
	}

	// Create a new VM and load the program
	vm := &interpreter.VM{}
	if err := vm.Load(root); err != nil {
		panic(err)
	}

	// Run the program
	if err := vm.Run(); err != nil {
		panic(err)
	}
}
