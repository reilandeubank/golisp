package interpreter

import (
	"fmt"
	"golisp/pkg/parser"
)

type Interpreter struct {
	environment *Environment
	globals     *Environment
}

// NewInterpreter defines an interpreter instance where the environment and globals are the same environment
func NewInterpreter() Interpreter {
	global := NewEnvironment()
	return Interpreter{environment: &global, globals: &global}
}

// evaluate calls the Accept method on a single expression
func (i *Interpreter) evaluate(expr parser.Expression) (interface{}, error) {
	return expr.Accept(i)
}

// Interpret will evaluate all expressions in the source code, printing out returned values
func (i *Interpreter) Interpret(exprs []parser.Expression) error {
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

// evaluateFunction will call evaluate the function's expression
// and then return the current environment to normal after completion
func (i *Interpreter) evaluateFunction(expression parser.Expression, environment Environment) (interface{}, error) {
	previous := i.environment

	defer func() {
		i.environment = previous
	}()

	i.environment = &environment
	return i.evaluate(expression)
}
