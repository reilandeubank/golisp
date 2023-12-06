package interpreter

import (
	"github.com/reilandeubank/golisp/pkg/scanner"
)

type environment struct {
	values map[string]interface{}
}

func NewEnvironment() environment {
	return environment{values: make(map[string]interface{})}
}

func (e *environment) get(name scanner.Token) (interface{}, error) {
	value, ok := e.values[name.Lexeme]
	if !ok {
		return nil, &RuntimeError{Token: name, Message: "Undefined variable '" + name.Lexeme + "'."}
	}
	return value, nil
}

func (e *environment) define(name string, value interface{}) {
	e.values[name] = value
}