package interpreter

import (
	// "fmt"
	"github.com/reilandeubank/golisp/pkg/parser"
)

type Interpreter struct {

}

func NewInterpreter() Interpreter {
	return Interpreter{}
}

func (i *Interpreter) evaluate(expr parser.Expression) (interface{}, error) {
	return expr.Accept(i)
}

func (i *Interpreter) Interpret(expr parser.Expression) (interface{}, error) {
	return i.evaluate(expr)
}