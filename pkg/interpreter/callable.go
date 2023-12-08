package interpreter

import (
	// "time"
	// "fmt"
	"github.com/reilandeubank/golisp/pkg/parser"
)

type LispCallable interface {
	Arity() int
	Call(i *Interpreter, arguments []interface{}) (interface{}, error)
	String() string
}

type LispFunction struct {
	Declaration parser.FuncDefinition
	Closure     *environment
	IsInitializer bool
}

func (l LispFunction) String() string {
	return "<fn " + l.Declaration.Name.Lexeme + ">"
}

func (l LispFunction) Arity() int {
	return len(l.Declaration.Params)
}

func (l LispFunction) Call(i *Interpreter, arguments []interface{}) (interface{}, error) {
	env := NewEnvironmentWithEnclosing(*l.Closure)

	for j, param := range l.Declaration.Params {
		env.define(param.Lexeme, arguments[j])
	}

	return i.evaluateFunction(l.Declaration.Body, env)
}