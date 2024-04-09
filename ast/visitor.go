// pkg/ast/visitor.go
package ast

// BasicVisitor is a basic implementation of the Visitor interface.
type BasicVisitor struct {
}

func (v *BasicVisitor) VisitExpr(e Expr) {
	// Implement the logic for visiting expression nodes
}

func (v *BasicVisitor) VisitStmt(s Stmt) {
	// Implement the logic for visiting statement nodes
}
