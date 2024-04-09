// pkg/stdlib/io.go
package stdlib

import (
	"fmt"
	"my-language/types"
)

// Print prints the given values to the console.
func Print(args ...interface{}) types.Type {
	for _, arg := range args {
		fmt.Print(arg)
	}
	fmt.Println()
	return types.NewPrimitiveType(types.Void)
}

// Println prints the given values to the console with a newline.
func Println(args ...interface{}) types.Type {
	fmt.Println(args...)
	return types.NewPrimitiveType(types.Void)
}
