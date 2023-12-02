package parser

import (
	"github.com/reilandeubank/golisp/pkg/scanner"
)

type Stmt interface {
	Accept(v StmtVisitor) (interface{}, error)
}

type ExprStmt struct {
	Expression Expression
}

func (e ExprStmt) Accept(v StmtVisitor) (interface{}, error) {
	return v.VisitExprStmt(e)
}

type SetStmt struct {
	Name        scanner.Token
	Initializer Expression
}

func (v SetStmt) Accept(visitor StmtVisitor) (interface{}, error) {
	return visitor.VisitSetStmt(v)
}

type FunctionStmt struct {
	Name        scanner.Token
	Params      []scanner.Token
	Body        []Stmt
}

func (f FunctionStmt) Accept(visitor StmtVisitor) (interface{}, error) {
	return visitor.VisitFunctionStmt(f)
}