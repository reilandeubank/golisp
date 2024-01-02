package parser

import (
	"golisp/pkg/scanner"
)

// Expression interface

// Expression is an interface that all expressions will implement
type Expression interface {
	Accept(v ExprVisitor) (interface{}, error)
	String() string
}

// Atom Interface

// Atom is a struct that implements the Expression interface
type Atom struct {
	Value interface{}
	Type  scanner.TokenType
}

// Accept is a method that visits the Atom expression and returns the result
func (l Atom) Accept(v ExprVisitor) (interface{}, error) {
	return v.VisitAtomExpr(l)
}

func (l Atom) String() string {
	return stringify(l.Value)
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
	Args    []Expression
}

func (k Keyword) Accept(v ExprVisitor) (interface{}, error) {
	return v.VisitKeywordExpr(k)
}

func (k Keyword) String() string {
	return scanner.KeywordsReverse[k.Keyword.Type]
}

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

type FuncDefinition struct {
	Name   scanner.Token
	Params []scanner.Token
	Body   Expression
}

func (f FuncDefinition) Accept(v ExprVisitor) (interface{}, error) {
	return v.VisitFuncDefinitionExpr(f)
}

func (f FuncDefinition) String() string {
	return "Define " + f.Name.Lexeme + " " + stringify(f.Params) + " " + f.Body.String()
}

// Call

// Call is a struct that implements the Expression interface
type Call struct {
	Callee   Expression
	Token    scanner.Token
	ArgsList []Expression
}

// Accept is a method that returns a string representation of the expression
func (c Call) Accept(v ExprVisitor) (interface{}, error) {
	return v.VisitCallExpr(c)
}

func (c Call) String() string {
	return c.Callee.String()
}

type Function struct {
	Name   scanner.Token
	Params []scanner.Token
	Body   Expression
}
