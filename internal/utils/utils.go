package utils

import (
	"fmt"
	"io/ioutil"
	"my-language/ast"
	"path/filepath"
)

// ReadFile reads the contents of a file.
func ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// JoinPath joins the given path components.
func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}

// Error represents an error with a position.
type Error struct {
	Pos     ast.Position
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d:%d: %s", e.Pos.Line, e.Pos.Column, e.Message)
}
