package interpreter

import (
	"fmt"
	"github.com/reilandeubank/golisp/pkg/parser"
)

type Interpreter struct {
	environment *environment
	globals    *environment
}

func NewInterpreter() Interpreter {
	global := NewEnvironment()
	return Interpreter{environment: &global, globals: &global}
}

func (i *Interpreter) evaluate(expr parser.Expression) (interface{}, error) {
	return expr.Accept(i)
}

func (i *Interpreter) Interpret(exprs []parser.Expression) (error) {
	for _, expr := range exprs {
		out, err := i.evaluate(expr)
		if out != nil {
			fmt.Println(out)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Interpreter) evaluateFunction(expression parser.Expression, environment environment) (interface{}, error) {
	previous := i.environment

	defer func() {
		i.environment = previous
	}()

	i.environment = &environment
	return i.evaluate(expression)
}