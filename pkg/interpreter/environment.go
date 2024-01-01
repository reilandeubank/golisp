package interpreter

import (
	"github.com/reilandeubank/golisp/pkg/scanner"
	"os"
)

// Environment allows for variable scope and closures. Less necessary in this implementation of Lisp
type Environment struct {
	enclosing *Environment
	values    map[string]interface{}
}

func NewEnvironment() Environment {
	return Environment{enclosing: nil, values: make(map[string]interface{})}
}

func NewEnvironmentWithEnclosing(Enclosing Environment) Environment {
	return Environment{enclosing: &Enclosing, values: make(map[string]interface{})}
}

// define a variable name as the passed value. only allowed in global scope
func (e *Environment) define(name string, value interface{}) {
	e.values[name] = value // this allows for variable redefinition. May be weird in normal code, but is useful for REPL
}

// get the value of a variable name. searches in enclosing environments, but throws if the variable is never found
func (e *Environment) get(name scanner.Token) (interface{}, error) {
	value, ok := e.values[name.Lexeme]
	if !ok && e.enclosing != nil {
		return e.enclosing.get(name)
	} else if name.Lexeme == "exit" {
		os.Exit(0)
		return value, nil
	} else if !ok {
		return nil, &RuntimeError{Token: name, Message: "Undefined variable '" + name.Lexeme + "'."}
	}
	return value, nil
}
