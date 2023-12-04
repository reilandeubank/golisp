package parser

import (
	"github.com/reilandeubank/golisp/pkg/scanner"
)

// Expression interface

// Interface in go is similar to an abstract class in Java
// Expression is an interface that all expressions will implement
type Expression interface {
	Accept(v ExprVisitor) (interface{}, error)
	String() string
}

// Literal

// Literal is a struct that implements the Expression interface
type Atom struct {
	Value interface{}
	Type  scanner.TokenType
}

// Accept() is a method that returns a string representation of the expression
func (l Atom) Accept(v ExprVisitor) (interface{}, error) {
	return v.VisitAtomExpr(l)
}

func (l Atom) String() string {
	return l.Value.(string)
}

type Operator struct {
	Operator scanner.Token
	Operands []Expression
}

func (o Operator) Accept(v ExprVisitor) (interface{}, error) {
	return v.VisitOperatorExpr(o)
}

func (o Operator) String() string {
	return o.Operator.Lexeme
}

// S-Expression
type ListExpr struct {
	Head Expression
	Tail []Expression
}

func (l ListExpr) Accept(v ExprVisitor) (interface{}, error) {
	return v.VisitListExpr(l)
}

func (l ListExpr) String() string {
	output := "(" + l.Head.String()
	for _, expr := range l.Tail {
		if expr != nil {
			// fmt.Println(reflect.TypeOf(expr))
			output += " " + expr.String()
		}
	}
	output += ")"
	return output
}

type Keyword struct {
	Keyword scanner.Token
}

func (k Keyword) Accept(v ExprVisitor) (interface{}, error) {
	return v.VisitKeywordExpr(k)
}

func (k Keyword) String() string {
	return k.Keyword.Lexeme
}

// // Unary

// // Unary is a struct that implements the Expression interface
// type Unary struct {
// 	Operator scanner.Token
// 	Right    Expression
// }

// // Accept() is a method that returns a string representation of the expression
// func (u Unary) Accept(v ExprVisitor) (interface{}, error) {
// 	return v.VisitUnaryExpr(u)
// }

// func (u Unary) String() string {
// 	return "(" + u.Operator.Lexeme + u.Right.String() + ")"
// }

// // Binary

// // Binary is a struct that implements the Expression interface
// type Binary struct {
// 	Left     Expression
// 	Operator scanner.Token
// 	Right    Expression
// }

// // Accept() is a method that returns a string representation of the expression
// func (b Binary) Accept(v ExprVisitor) (interface{}, error) {
// 	return v.VisitBinaryExpr(b)
// }

// func (b Binary) String() string {
// 	return "(" + b.Operator.Lexeme + b.Left.String() + b.Right.String() + ")"
// }

// Variable

// Variable is a struct that implements the Expression interface
type Symbol struct {
	Name scanner.Token
}

// Accept() is a method that returns a string representation of the expression
func (s Symbol) Accept(v ExprVisitor) (interface{}, error) {
	return v.VisitSymbolExpr(s)
}

func (s Symbol) String() string {
	return s.Name.Lexeme
}

// // Assignment

// // Assignment is a struct that implements the Expression interface
// type Assign struct {
// 	Name  scanner.Token
// 	Value Expression
// }

// // Accept() is a method that returns a string representation of the expression
// func (a Assign) Accept(v ExprVisitor) (interface{}, error) {
// 	return v.VisitAssignExpr(a)
// }

// func (a Assign) String() string {
// 	return " set " + a.Name.Lexeme + a.Value.String()
// }

// // Logical

// // Logical is a struct that implements the Expression interface
// type Logical struct {
// 	Left     Expression
// 	Operator scanner.Token
// 	Right    Expression
// }

// // Accept() is a method that returns a string representation of the expression
// func (l Logical) Accept(v ExprVisitor) (interface{}, error) {
// 	return v.VisitLogicalExpr(l)
// }

// func (l Logical) String() string {
// 	return "(" + l.Operator.Lexeme + l.Left.String() + l.Right.String() + ")"
// }

// // Call

// // Call is a struct that implements the Expression interface
// type Call struct {
// 	Callee    Expression
// 	Paren     scanner.Token
// 	Arguments []Expression
// }

// // Accept() is a method that returns a string representation of the expression
// func (c Call) Accept(v ExprVisitor) (interface{}, error) {
// 	return v.VisitCallExpr(c)
// }

// func (c Call) String() string {
// 	return c.Callee.String()
// }

// // Cond
// type Cond struct {
// 	Condition  Expression
// 	ThenBranch Expression
// 	ElseBranch Expression
// }

// func (c Cond) Accept(v ExprVisitor) (interface{}, error) {
// 	return v.VisitCondExpr(c)
// }

// func (c Cond) String() string {
// 	return "(cond " + c.Condition.String() + " " + c.ThenBranch.String() + " " + c.ElseBranch.String() + ")"
// }
