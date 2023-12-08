package interpreter

import (
	"fmt"
	"reflect"

	"github.com/reilandeubank/golisp/pkg/parser"
	"github.com/reilandeubank/golisp/pkg/scanner"
)

func (i *Interpreter) VisitListExpr(l parser.ListExpr) (interface{}, error) {
	switch head := l.Head.(type) {
	case parser.Operator:
		head.Operands = l.Tail
		result, err := i.evaluate(head)
		if err != nil {
			return nil, err
		}
		if result == false {
			return nil, nil
		}
		return result, nil
	case parser.Keyword:
		head.Args = l.Tail
		return i.evaluate(head)
	case parser.Atom:
		returnList := parser.ListExpr{Head: head, Tail: l.Tail}
		return returnList, nil
	}

	return nil, fmt.Errorf("LISTEXPR not implemented")
}

func (i *Interpreter) VisitKeywordExpr(k parser.Keyword) (interface{}, error) {
	switch k.Keyword.Type {
	case scanner.TRUE:
		return true, nil
	// case scanner.FALSE:
	// 	return false, nil
	case scanner.NIL:
		return nil, nil
	case scanner.CAR:
		car, err := i.evaluate(k.Args[0])
		switch car.(type) {
		case parser.ListExpr:
			car = car.(parser.ListExpr).Head
			return i.evaluate(car.(parser.Expression))
		}
		if err != nil {
			return nil, err
		}
		return car, nil
	case scanner.CDR:
		output, err := i.evaluate(k.Args[0])
		if err != nil {
			return nil, err
		}
		list, ok := output.(parser.ListExpr)
		if !ok {
			return nil, &RuntimeError{Token: k.Keyword, Message: "CDR operation must have a list as the first operand"}
		}

		if len(list.Tail) > 1 {
			return parser.ListExpr{Head: list.Tail[0], Tail: list.Tail[1:]}, nil
		} else {
			return parser.ListExpr{Head: list.Tail[0]}, nil
		}
	case scanner.COND:
		for j := 0; j < len(k.Args); j += 2 {
			condition, err := i.evaluate(k.Args[j])
			if err != nil {
				return nil, err
			}
			if isTruthy(condition) && j+1 < len(k.Args) {
				return i.evaluate(k.Args[j+1])
			}
		}
		return nil, &RuntimeError{Token: k.Keyword, Message: "Lack of true condition"}
	case scanner.NUMBERQ:
		if len(k.Args) != 1 {
			return nil, &RuntimeError{Token: k.Keyword, Message: "NUMBER? operation must have 1 operand"}
		}
		expr, err := i.evaluate(k.Args[0])
		if err != nil {
			return nil, err
		}
		return checkNumberOperand(k.Keyword, expr)
	case scanner.SYMBOLQ:
		if len(k.Args) != 1 {
			return nil, &RuntimeError{Token: k.Keyword, Message: "SYMBOL? operation must have 1 operand"}
		}
		return reflect.TypeOf(k.Args[0]) == reflect.TypeOf(parser.Symbol{}), nil
	case scanner.LISTQ:
		if len(k.Args) != 1 {
			return nil, &RuntimeError{Token: k.Keyword, Message: "LIST? operation must have 1 operand"}
		}
		expr, err := i.evaluate(k.Args[0])
		if err != nil {
			return nil, err
		}
		return reflect.TypeOf(expr) == reflect.TypeOf(parser.ListExpr{}), nil
	case scanner.NILQ:
		if len(k.Args) != 1 {
			return nil, &RuntimeError{Token: k.Keyword, Message: "NIL? operation must have 1 operand"}
		}
		expr, err := i.evaluate(k.Args[0])
		if err != nil {
			return nil, err
		}
		return expr == nil, nil
	case scanner.ANDQ:
		if len(k.Args) != 2 {
			return nil, &RuntimeError{Token: k.Keyword, Message: "AND? operation must have 2 operands"}
		}
		left, err := i.evaluate(k.Args[0])
		if err != nil {
			return nil, err
		}
		right, err := i.evaluate(k.Args[1])
		if err != nil {
			return nil, err
		}
		return isTruthy(left) && isTruthy(right), nil
	case scanner.ORQ:
		if len(k.Args) != 2 {
			return nil, &RuntimeError{Token: k.Keyword, Message: "OR? operation must have 2 operands"}
		}
		left, err := i.evaluate(k.Args[0])
		if err != nil {
			return nil, err
		}
		right, err := i.evaluate(k.Args[1])
		if err != nil {
			return nil, err
		}
		result := isTruthy(left) || isTruthy(right)
		if result == false {
			return nil, nil
		}
		return result, nil
	case scanner.NOTQ:
		if len(k.Args) != 2 {
			return nil, &RuntimeError{Token: k.Keyword, Message: "NOT? operation must have 1 operand"}
		}
		expr, err := i.evaluate(k.Args[0])
		if err != nil {
			return nil, err
		}
		return !isTruthy(expr), nil
	case scanner.SET:
		if len(k.Args) != 2 {
			return nil, &RuntimeError{Token: k.Keyword, Message: "SET operation must have 2 operands"}
		}
		if reflect.TypeOf(k.Args[0]) != reflect.TypeOf(parser.Symbol{}) {
			return nil, &RuntimeError{Token: k.Keyword, Message: "SET operation must have a symbol as the first operand"}
		}
		value, err := i.evaluate(k.Args[1])
		if err != nil {
			return nil, err
		}
		i.environment.define(k.Args[0].(parser.Symbol).Name.Lexeme, value)
		return nil, nil
	}
	return nil, fmt.Errorf("KEYWORDEXPR not implemented")
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
	return nil, &RuntimeError{Token: o.Operator, Message: "Invalid operator"}
}

func (i *Interpreter) VisitAtomExpr(a parser.Atom) (interface{}, error) {
	return a.Value, nil
}

func (i *Interpreter) VisitSymbolExpr(s parser.Symbol) (interface{}, error) {
	return i.environment.get(s.Name)
}

func (i *Interpreter) VisitCallExpr (c parser.Call) (interface{}, error) {
	callee, err := i.evaluate(c.Callee)
	if err != nil {
		return nil, err
	}

	arguments := make([]interface{}, len(c.ArgsList))
	for j, argument := range c.ArgsList {
		arguments[j], err = i.evaluate(argument)
		if err != nil {
			return nil, err
		}
	}

	function, ok := callee.(LispCallable)
	if !ok {
		return nil, &RuntimeError{Token: c.Token, Message: "Can only call functions."}
	}
	if len(arguments) != function.Arity() {
		return nil, &RuntimeError{Token: c.Token, Message: "Expected " + fmt.Sprint(function.Arity()) + " arguments but got " + fmt.Sprint(len(arguments)) + "."}
	}

	return function.Call(i, arguments)
}

func (i *Interpreter) VisitFuncDefinitionExpr(f parser.FuncDefinition) (interface{}, error) {
	function := LispFunction{Declaration: f, Closure: i.environment, IsInitializer: false}
	i.environment.define(f.Name.Lexeme, function)
	return nil, nil
}
