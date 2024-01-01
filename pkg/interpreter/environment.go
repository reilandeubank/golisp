package interpreter

import (
	"github.com/reilandeubank/golisp/pkg/scanner"
)

// environment allows for variable scope and closures. Less necessary in this implementation of Lisp
type environment struct {
	enclosing *environment
	values    map[string]interface{}
}

func NewEnvironment() environment {
	return environment{enclosing: nil, values: make(map[string]interface{})}
}

func NewEnvironmentWithEnclosing(Enclosing environment) environment {
	return environment{enclosing: &Enclosing, values: make(map[string]interface{})}
}

// define a variable name as the passed value. only allowed in global scope
func (e *environment) define(name string, value interface{}) {
	e.values[name] = value // this allows for variable redefinition. May be weird in normal code, but is useful for REPL
}

// get the value of a variable name. searches in enclosing environments, but throws if the variable is never found
func (e *environment) get(name scanner.Token) (interface{}, error) {
	value, ok := e.values[name.Lexeme]
	if !ok && e.enclosing != nil {
		return e.enclosing.get(name)
	} else if !ok {
		return nil, &RuntimeError{Token: name, Message: "Undefined variable '" + name.Lexeme + "'."}
	}
	return value, nil
}
