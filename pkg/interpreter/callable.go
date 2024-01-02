package interpreter

import (
	// "time"
	// "fmt"
	"golisp/pkg/parser"
)

// LispCallable  is the interface that LispFunction implements (unnecessary?)
type LispCallable interface {
	Arity() int
	Call(i *Interpreter, arguments []interface{}) (interface{}, error)
	String() string
}

// LispFunction is the struct that all functions in the language are stored as
type LispFunction struct {
	Declaration   parser.FuncDefinition
	Closure       *Environment
	IsInitializer bool
}

// String returns a string representation of the function for debugging purposes
func (l LispFunction) String() string {
	return "<fn " + l.Declaration.Name.Lexeme + ">"
}

// Arity is simply the number of parameters for a function (arity must match for function calls)
func (l LispFunction) Arity() int {
	return len(l.Declaration.Params)
}

// Call generates a new environment, defines each parameter as the passed arguments, then evaluates
func (l LispFunction) Call(i *Interpreter, arguments []interface{}) (interface{}, error) {
	env := NewEnvironmentWithEnclosing(*l.Closure)

	for j, param := range l.Declaration.Params {
		env.define(param.Lexeme, arguments[j])
	}

	return i.evaluateFunction(l.Declaration.Body, env)
}
