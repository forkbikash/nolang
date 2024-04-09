// pkg/parser/scanner.go
package parser

import "my-language/ast"

// Token represents a lexical token.
type Token struct {
	Type     TokenType
	Literal  string
	Position ast.Position
}

// TokenType represents the type of a token.
type TokenType int

const (
	Illegal TokenType = iota
	EOF

	Ident  // Identifiers
	Int    // Integer literals
	Float  // Floating-point literals
	String // String literals
	Bool   // Boolean literals

	// Operators
	Plus
	Minus
	Mul
	Div
	Equal
	NotEqual
	LessThan
	GreaterThan
	Assign

	// Delimiters
	Comma
	Semicolon
	Colon
	LeftParen
	RightParen
	LeftBrace
	RightBrace
)

// Scan scans the input and returns the next token.
func Scan(input string) *Token {
	// Implement the scanner logic to convert the input into a sequence of tokens
	return nil
}
