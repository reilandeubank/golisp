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

func (l LispFunction) Call(i *Interpreter, arguments []interface{}) (retVal interface{}, errVal error) {
	env := NewEnvironmentWithEnclosing(*l.Closure)

	for j, param := range l.Declaration.Params {
		env.define(param.Lexeme, arguments[j])
	}

	// _, err := i.executeBlock(l.Declaration.Body, env)
	// if err != nil {
	// 	return nil, err
	// }

	// if l.IsInitializer {
	// 	return l.Closure.getAt(0, "this")
	// }

	return retVal, errVal
}

// type clock struct{}

// func (c *clock) Arity() int {
// 	return 0
// }

// func (c *clock) Call(i *Interpreter, arguments []interface{}) (interface{}, error) {
// 	return float64(time.Now().UnixMilli()) / 1000, nil
// }

// func (c clock) String() string {
// 	return "<native fn>"
// }

// type toStr struct{}

// func (t *toStr) Arity() int {
// 	return 1
// }

// func (t *toStr) Call(i *Interpreter, arguments []interface{}) (interface{}, error) {
// 	return fmt.Sprintf("%v", arguments[0]), nil
// }

// func (t toStr) String() string {
// 	return "<native fn>"
// }