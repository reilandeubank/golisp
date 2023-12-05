package interpreter

import (
	"fmt"
	"reflect"

	"github.com/reilandeubank/golisp/pkg/parser"
	"github.com/reilandeubank/golisp/pkg/scanner"
)

func (i *Interpreter) VisitListExpr(l parser.ListExpr) (interface{}, error) {
	//fmt.Println(l.Head.String())

	switch head := l.Head.(type) {
	case parser.Operator:
		// fmt.Println("OPERATOR", head.String())
		head.Operands = l.Tail
		return i.evaluate(head)
	case parser.Keyword:
		// fmt.Println("KEYWORD", head.String())
		head.Args = l.Tail
		return i.evaluate(head)
	case parser.Atom:
		// fmt.Println("ATOM LIST", head.String())
		// returnList := prepend(head, l.Tail)
		// fmt.Println(returnList)
		returnList := parser.ListExpr{Head: head, Tail: l.Tail}
		return returnList, nil
    }

	return nil, fmt.Errorf("LISTEXPR not implemented")
}

func (i *Interpreter) VisitKeywordExpr(k parser.Keyword) (interface{}, error) {
	// fmt.Println()
	// fmt.Println(k.String())
	// fmt.Println()

	switch k.Keyword.Type {
	case scanner.TRUE:
		return true, nil
	// case scanner.FALSE:
	// 	return false, nil
	case scanner.NIL:
		return nil, nil
	case scanner.CAR:
		// fmt.Println("CAR of", k.Args)
		car, err := i.evaluate(k.Args[0])
		switch car.(type) {
		case parser.ListExpr:
			car = car.(parser.ListExpr).Head
			return i.evaluate(car.(parser.Expression))
		}
		// fmt.Println("Car is", car)
		if err != nil {
			return nil, err
		}
		return car, nil
	case scanner.CDR:
		return i.cdr(k)
	case scanner.COND:
		// fmt.Println("COND", k.Args)
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
		// fmt.Println(k.String())
		if len(k.Args) != 1 {
			return nil, &RuntimeError{Token: k.Keyword, Message: "NUMBER? operation must have 1 operand"}
		}
		expr, err := i.evaluate(k.Args[0])
		if err != nil {
			return nil, err
		}
		return checkNumberOperand(k.Keyword, expr)
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
			return nil, &RuntimeError{Token: k.Keyword, Message: "NUMBER? operation must have 1 operand"}
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
		return isTruthy(left) || isTruthy(right), nil
	case scanner.NOTQ:
		if len(k.Args) != 2 {
			return nil, &RuntimeError{Token: k.Keyword, Message: "NOT? operation must have 1 operand"}
		}
		expr, err := i.evaluate(k.Args[0])
		if err != nil {
			return nil, err
		}
		return !isTruthy(expr), nil
	}
	return nil, fmt.Errorf("KEYWORDEXPR not implemented")
}

func (i *Interpreter) VisitOperatorExpr(o parser.Operator) (interface{}, error) {
	if len(o.Operands) != 2 {
		fmt.Println(o.Operands)
		return nil, &RuntimeError{Token: o.Operator, Message: "Binary operation must only have two operands"}
	}
	left, err := i.evaluate(o.Operands[0])
	if err != nil {
		return nil, err
	}
	// fmt.Println("left", left)
	right, err := i.evaluate(o.Operands[1])
	if err != nil {
		return nil, err
	}
	// fmt.Println("right", right)

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
	return nil, fmt.Errorf("SYMBOLEXPR not implemented")
}
