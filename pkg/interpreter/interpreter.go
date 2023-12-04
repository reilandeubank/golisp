package interpreter

import (
	"fmt"
	"github.com/reilandeubank/golisp/parser"
)

type Interpreter struct {

}

func NewInterpreter() Interpreter {
	return Interpreter{}
}

func (i *Interpreter) evaluate(expr parser.Expression) (interface{}, error) {
	return expr.Accept(i)
}