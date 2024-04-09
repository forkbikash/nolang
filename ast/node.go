// pkg/ast/node.go
package ast

import "my-language/types"

// Node represents a node in the abstract syntax tree.
type Node interface {
	Pos() Position
	Type() types.Type
	Accept(Visitor)
}

// Position represents the position of a node in the source code.
type Position struct {
	Line   int
	Column int
}

// Expr represents an expression node.
type Expr Node

// Stmt represents a statement node.
type Stmt Node

// Visitor defines the interface for visiting AST nodes.
type Visitor interface {
	VisitExpr(Expr)
	VisitStmt(Stmt)
}
