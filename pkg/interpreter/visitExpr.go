package interpreter

import (
	"fmt"
	"reflect"

	"github.com/reilandeubank/golisp/pkg/parser"
	"github.com/reilandeubank/golisp/pkg/scanner"
)

func (i *Interpreter) VisitListExpr(l parser.ListExpr) (interface{}, error) {
	if o, ok := l.Head.(*parser.Operator); ok {
		o.Operands = l.Tail
		return i.evaluate(o)
	}
	return nil, fmt.Errorf("not implemented")
}

func (i *Interpreter) VisitKeywordExpr(k parser.Keyword) (interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}

func (i *Interpreter) VisitOperatorExpr(o parser.Operator) (interface{}, error) {
	if len(o.Operands) != 2 {
		return nil, &RuntimeError{Token: o.Operator, Message: "Binary operation must only have two operands"}
	}
	left, err := i.evaluate(o.Operands[0])
	if err != nil {
		return nil, err
	}
	right, err := i.evaluate(o.Operands[1])
	if err != nil {
		return nil, err
	}

	switch o.Operator.Type {
	case scanner.MINUS:
		err = checkNumberOperands(o.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) - right.(float64), nil
	case scanner.PLUS:
		if reflect.TypeOf(left) == reflect.TypeOf("") && reflect.TypeOf(right) == reflect.TypeOf("") {
			return left.(string) + right.(string), nil
		}
		if reflect.TypeOf(left) == reflect.TypeOf(0.0) && reflect.TypeOf(right) == reflect.TypeOf(0.0) {
			return left.(float64) + right.(float64), nil
		}
		return nil, err
	case scanner.SLASH:
		err = checkNumberOperands(o.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) / right.(float64), nil
	case scanner.STAR:
		err = checkNumberOperands(o.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) * right.(float64), nil
	case scanner.GREATER:
		err = checkNumberOperands(o.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) > right.(float64), nil
	case scanner.LESS:
		err = checkNumberOperands(o.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) < right.(float64), nil
	case scanner.EQUAL:
		return isEqual(left, right), nil
	}
	return nil, fmt.Errorf("not implemented")
}

func (i *Interpreter) VisitAtomExpr(a parser.Atom) (interface{}, error) {
	return a.Value, nil
}

func (i *Interpreter) VisitSymbolExpr(s parser.Symbol) (interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}
