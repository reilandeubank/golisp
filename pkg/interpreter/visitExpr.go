package interpreter

import (
	"fmt"

	"github.com/reilandeubank/golisp/parser"
)

func (i *Interpreter) VisitListExpr(l parser.ListExpr) (interface{}, error) {
	if isOperator(l.Head) {
		return i.VisitOperatorExpr(o)
	}
	return nil, fmt.Errorf("not implemented")
}

func (i *Interpreter) VisitKeywordExpr(k parser.Keyword) (interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}

func (i *Interpreter) VisitOperatorExpr(o parser.Operator) (interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}

func (i *Interpreter) VisitAtomExpr(a parser.Atom) (interface{}, error) {
	return a.Value, nil
}

func (i *Interpreter) VisitSymbolExpr(s parser.Symbol) (interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}
