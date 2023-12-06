package interpreter

import (
	// "fmt"
	"github.com/reilandeubank/golisp/pkg/parser"
)

type Interpreter struct {
	environment *environment
}

func NewInterpreter() Interpreter {
	global := NewEnvironment()
	return Interpreter{environment: &global}
}

func (i *Interpreter) evaluate(expr parser.Expression) (interface{}, error) {
	return expr.Accept(i)
}

func (i *Interpreter) Interpret(expr parser.Expression) (interface{}, error) {
	return i.evaluate(expr)
}